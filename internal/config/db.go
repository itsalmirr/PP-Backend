package config

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"backend.com/go-backend/ent"
)

type Database struct {
	Client *ent.Client
	pool   *pgxpool.Pool
}

func ConnectDatabase(ctx context.Context, cfg *Config) (*Database, error) {
	// Use pgx configuration for better connection control
	connConfig, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	))

	if err != nil {
		panic(fmt.Sprintf("Failed to parse connection string: %v", err))
	}

	// Configure connection pool
	connConfig.MaxConns = 50
	connConfig.MinConns = 10
	connConfig.MaxConnLifetime = time.Hour
	connConfig.MaxConnIdleTime = time.Minute * 30

	pool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to create connection pool: %v", err))
	}

	db := stdlib.OpenDB(*connConfig.ConnConfig)

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)

	if err := db.Ping(); err != nil {
		pool.Close() // Clean up the pool if the ping fails
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	driver := sql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(driver))

	fmt.Printf("Connected to database: %s", cfg.DBName)

	return &Database{
		Client: client,
		pool:   pool,
	}, nil
}

func (db *Database) Migrate(ctx context.Context) error {
	// Run migrations with timeout
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := db.Client.Schema.Create(
		ctx,
		// schema.WithAtlas(true),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	); err != nil {
		return fmt.Errorf("failed to run database migrations: %w", err)
	}
	return nil
}

func (db *Database) Close() error {
	err := db.Client.Close()
	db.pool.Close()
	return err
}

func SessionStorage(cfg *Config) redis.Store {
	// Validate secret
	secretHex := cfg.SessionKey
	if len(secretHex) != 64 && len(secretHex) != 128 { // Check byte length
		panic("SESSION_SECRET must be 32 or 64 bytes (64/128 hex chars)")
	}

	key, err := hex.DecodeString(secretHex)
	if err != nil {
		panic("Failed to decode SESSION_SECRET: " + err.Error())
	}

	store, err := redis.NewStore(
		10,
		"tcp",
		cfg.RedisURL,
		"",
		key,
		key,
	)

	if err != nil {
		panic("Failed to create Redis store: " + err.Error())
	}

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3 * 24 * 60 * 60,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return store
}

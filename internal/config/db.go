package config

import (
	"context"
	"encoding/hex"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"ppgroup.ppgroup.com/ent"
)

const (
	DefaultMaxConns        = 50
	DefaultMinConns        = 10
	DefaultConnMaxLifetime = time.Hour
	DefaultConnMaxIdleTime = 30 * time.Minute
)

type Database struct {
	Client *ent.Client
	pool   *pgxpool.Pool
}

func ConnectDatabase(ctx context.Context, cfg *Config) (*Database, error) {
	connConfig, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	))
	if err != nil {
		return nil, fmt.Errorf("parsing connection string: %w", err)
	}

	connConfig.MaxConns = DefaultMaxConns
	connConfig.MinConns = DefaultMinConns
	connConfig.MaxConnLifetime = DefaultConnMaxLifetime
	connConfig.MaxConnIdleTime = DefaultConnMaxIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	db := stdlib.OpenDB(*connConfig.ConnConfig)

	db.SetMaxOpenConns(int(DefaultMaxConns))
	db.SetMaxIdleConns(int(DefaultMinConns))
	db.SetConnMaxLifetime(DefaultConnMaxLifetime)
	db.SetConnMaxIdleTime(DefaultConnMaxIdleTime)

	if err := db.Ping(); err != nil {
		pool.Close()
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	driver := sql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver))

	slog.Info("connected to database", "name", cfg.DBName)

	return &Database{
		Client: client,
		pool:   pool,
	}, nil
}

func (db *Database) Migrate(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := db.Client.Schema.Create(
		ctx,
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	); err != nil {
		return fmt.Errorf("running database migrations: %w", err)
	}
	return nil
}

func (db *Database) Close() error {
	err := db.Client.Close()
	db.pool.Close()
	return err
}

func SessionStorage(cfg *Config) (redis.Store, error) {
	secretHex := cfg.SessionKey
	if len(secretHex) != 64 && len(secretHex) != 128 {
		return nil, fmt.Errorf("SESSION_KEY must be 32 or 64 bytes (64/128 hex chars), got %d chars", len(secretHex))
	}

	key, err := hex.DecodeString(secretHex)
	if err != nil {
		return nil, fmt.Errorf("decoding SESSION_KEY: %w", err)
	}

	store, err := redis.NewStore(
		10,
		"tcp",
		cfg.RedisURL,
		"",
		"",
		key,
		key,
	)
	if err != nil {
		return nil, fmt.Errorf("creating Redis store: %w", err)
	}

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3 * 24 * 60 * 60,
		HttpOnly: true,
		Secure:   cfg.CookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
	return store, nil
}

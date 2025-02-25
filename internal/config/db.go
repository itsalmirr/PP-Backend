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

// var DB *gorm.DB

// func ConnectDatabase(cfg *Config) {
// 	dsn := fmt.Sprintf(
// 		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
// 		cfg.DBHost,
// 		cfg.DBUser,
// 		cfg.DBPassword,
// 		cfg.DBName,
// 		cfg.DBPort,
// 	)

// 	// Gorm dafabase config
// 	config := &gorm.Config{
// 		PrepareStmt:            true,
// 		Logger:                 logger.Default.LogMode(logger.Info),
// 		SkipDefaultTransaction: true,
// 		NowFunc: func() time.Time {
// 			return time.Now().UTC() // for consistent timezone
// 		},
// 	}

// 	database, err := gorm.Open(postgres.New(postgres.Config{
// 		DriverName: "pgx",
// 		DSN:        dsn,
// 	}), config)
// 	if err != nil {
// 		panic("Failed to connect to database!")
// 	}

// 	sqlDB, err := database.DB()
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to get DB instance: %v", err))
// 	}
// 	sqlDB.SetMaxIdleConns(10)
// 	sqlDB.SetMaxOpenConns(100)
// 	sqlDB.SetConnMaxLifetime(time.Hour)

// 	// Migrations
// 	err = database.AutoMigrate(
// 		&models.User{},
// 		&models.Listing{},
// 		&models.Realtor{},
// 	)
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to migrate: %v", err))
// 	}

// 	fmt.Println("Database connected and migrated successfully!")
// 	DB = database // Assign the database connection to the global variable
// }

func SessionStorage(cfg *Config) redis.Store {
	// Validate secret
	secretHex := cfg.SessionSecret
	if len(secretHex) != 32 && len(secretHex) != 64 { // Check byte length
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
		key, // Pass to both key arguments
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

package config

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"backend.com/go-backend/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Client *ent.Client

func Connectdatabase(cfg *Config) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// Create a new Ent client
	drv, err := sql.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("Failed to create database driver: %v", err)
		panic("Failed to create database driver")
	}

	// Configure connection pool settings
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	// Create the Ent client
	Client = ent.NewClient(ent.Driver(drv))

	// Run migrations using Atlas
	ctx := context.Background()
	if err := migrateDatabase(ctx, Client); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	fmt.Println("Database connected and migrated successfully!")
}

func migrateDatabase(ctx context.Context, client *ent.Client) error {
	// Create an Atlas migration engine
	opts := []schema.MigrateOption{
		schema.WithGlobalUniqueID(true),
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
	}

	// Create a migration plan
	err := client.Schema.Create(ctx, opts...)
	if err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	return nil
}

// CloseDatabase closes the database connection
func CloseDatabase() {
	if Client != nil {
		Client.Close()
	}
}

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

package migration

import (
	"context"
	"fmt"

	"github.com/chanhlab/go-kit/logger"
	"github.com/chanhlab/golang-service-example/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigration runs database migrations using golang-migrate
func RunMigration(_ context.Context) error {
	fmt.Printf("Running migrations \n")
	config.NewConfig()
	cfg := config.AppConfig
	
	// Convert string level to int for logger compatibility
	logLevel := 0
	switch cfg.Logging.Level {
	case "debug":
		logLevel = -1
	case "info":
		logLevel = 0
	case "warn":
		logLevel = 1
	case "error":
		logLevel = 2
	}
	
	logger.Init(int8(logLevel), cfg.Logging.TimeFormat)
	
	// Build database URL
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	
	// Create migrate instance
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()
	
	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	
	if err == migrate.ErrNoChange {
		fmt.Println("No migrations to run")
	} else {
		fmt.Println("Migrations completed successfully")
	}
	
	return nil
}

// RunMigrationDown rolls back the last migration
func RunMigrationDown(_ context.Context) error {
	fmt.Printf("Rolling back migrations \n")
	config.NewConfig()
	cfg := config.AppConfig
	
	// Build database URL
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	
	// Create migrate instance
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()
	
	// Run one step down
	if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to rollback migration: %w", err)
	}
	
	if err == migrate.ErrNoChange {
		fmt.Println("No migrations to rollback")
	} else {
		fmt.Println("Migration rollback completed successfully")
	}
	
	return nil
}

// GetMigrationVersion returns the current migration version
func GetMigrationVersion(_ context.Context) error {
	config.NewConfig()
	cfg := config.AppConfig
	
	// Build database URL
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	
	// Create migrate instance
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()
	
	version, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			fmt.Println("No migrations have been run")
			return nil
		}
		return fmt.Errorf("failed to get migration version: %w", err)
	}
	
	fmt.Printf("Current migration version: %d (dirty: %v)\n", version, dirty)
	return nil
}

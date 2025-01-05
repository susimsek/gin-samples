package config

import (
	"database/sql"
	"errors"
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file" // File source driver
)

var DatabaseConfig DatabaseInitializer = &RealDatabaseConfig{}

// DatabaseInitializer interface for initializing the database
type DatabaseInitializer interface {
	InitDB() *gorm.DB
}

// RealDatabaseConfig is the production implementation
type RealDatabaseConfig struct{}

func (r *RealDatabaseConfig) InitDB() *gorm.DB {
	sqlDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}

	runMigrations(sqlDB)

	gormDB, err := gorm.Open(sqlite.New(sqlite.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize GORM: %v", err)
	}

	return gormDB
}

func runMigrations(db *sql.DB) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create SQLite driver: %v", err)
	}

	migrationPath := filepath.Join("resources", "db", "migrations")
	m, err := migrate.NewWithDatabaseInstance("file://"+migrationPath, "sqlite3", driver)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

package config

import (
	"database/sql"
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file" // File source driver
)

func InitDB() *gorm.DB {
	// SQLite in-memory database with shared cache
	sqlDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}

	// Run migrations
	runMigrations(sqlDB)

	// Use GORM with the same database connection
	gormDB, err := gorm.Open(sqlite.New(sqlite.Config{
		Conn: sqlDB, // Share the same connection
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

	m, err := migrate.NewWithDatabaseInstance(
		"file://resources/db/migrations", // Migration files path
		"sqlite3",                        // Database driver name
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	// Apply all up migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

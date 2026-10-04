package postgres

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"github.com/gorkagg10/equity-calculator-api/internal/config"
)

//go:embed migrations/*.sql
var fs embed.FS

func Migrate(pgClient *sql.DB, databaseName, migrationsPath string) error {
	migration, err := newMigrator(pgClient, databaseName, migrationsPath)
	if err != nil {
		return fmt.Errorf("loading migration files: %w", err)
	}
	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	if errors.Is(err, migrate.ErrNoChange) {
		slog.Info("migration", slog.String("error", err.Error()))
	}
	return nil
}

func NewDatabaseClient(ctx context.Context, databaseConfig *config.DatabaseConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		databaseConfig.Host, databaseConfig.Port, databaseConfig.Username, databaseConfig.Password,
		databaseConfig.Database, databaseConfig.SSLMode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("opening database connection: %w", err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("pinging database connection: %w", err)
	}
	return db, nil
}

func newMigrator(pgClient *sql.DB, databaseName, migrationsPath string) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(pgClient, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("create pg driver: %w", err)
	}
	d, err := iofs.New(fs, migrationsPath)
	if err != nil {
		return nil, fmt.Errorf("read migrations: %w", err)
	}
	m, err := migrate.NewWithInstance("iofs", d, databaseName, driver)
	if err != nil {
		return nil, fmt.Errorf("create migrate instance: %w", err)
	}

	return m, nil
}

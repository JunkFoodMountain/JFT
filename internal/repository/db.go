package repository

import (
	"database/sql"
	"embed"
	"fmt"
	"time"

	_ "github.com/lib/pq" // need this for the postgres driver
	"github.com/pressly/goose/v3"
)

//go:generate go run ./cmd/generate_database.go
//go:generate sqlboiler --wipe psql
//go:embed migrations/*.sql
var embedMigrations embed.FS

// MigrateDB Migrates the database
func MigrateDB(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	return goose.Up(db, "migrations")
}

func OpenDB(conf Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("sslmode=disable host=%s user=%s password=%s", conf.Host, conf.Username, conf.Password)

	db, err := connectWithRetries(connectionString, conf.Retries, time.Second)
	if err != nil {
		return nil, err
	}

	_, _ = db.Exec("CREATE DATABASE jft")

	connectionString += " dbname=jft"
	return connectWithRetries(connectionString, conf.Retries, time.Second)
}

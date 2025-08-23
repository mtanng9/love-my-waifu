package db

import (
	"context"
	"database/sql"
	"embed"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type DB struct {
	Conn *sql.DB
}

//go:embed migrations/*.sql
var embedMigration embed.FS

func ConnectDB(db *DB) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgresql://postgres:lovemywaifu@0.0.0.0:5469/postgres?connect_timeout=10&application_name=lovemywaifu")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	connector := stdlib.GetPoolConnector(pool)
	db.Conn = sql.OpenDB(connector)
}

func MigrateDB(db *DB) {
	goose.SetBaseFS(embedMigration)
	err := goose.SetDialect(string(goose.DialectPostgres))
	if err != nil {
		log.Fatalf("could not set database dialect to postgres: %v", err)
	}

	err = goose.Up(db.Conn, "migrations")
	if err != nil {
		log.Fatalf("could not migrate db: %v", err)
	}
}

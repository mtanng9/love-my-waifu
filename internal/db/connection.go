package db

import (
	"context"
	"database/sql"
	"embed"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigration embed.FS

func ConnectDB() (*pgx.Conn, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgresql://postgres:lovemywaifu@0.0.0.0:5469/postgres?connect_timeout=10&application_name=lovemywaifu")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func MigrateDB() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgresql://postgres:lovemywaifu@0.0.0.0:5469/postgres?connect_timeout=10&application_name=lovemywaifu")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	connector := stdlib.GetPoolConnector(pool)
	db := sql.OpenDB(connector)

	goose.SetBaseFS(embedMigration)
	err = goose.SetDialect(string(goose.DialectPostgres))
	if err != nil {
		log.Fatalf("could not set database dialect to postgres: %v", err)
	}

	err = goose.Up(db, "migrations")
	if err != nil {
		log.Fatalf("could not migrate db: %v", err)
	}
}

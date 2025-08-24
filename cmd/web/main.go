package main

import (
	"log/slog"
	"os"

	"github.com/mtanng9/love-my-waifu/internal/db"
	"github.com/mtanng9/love-my-waifu/internal/server"
)

func main() {
	logHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(logHandler)
	errorLog := slog.NewLogLogger(logHandler, slog.LevelError)

	database, err := db.ConnectDB()
	if err != nil {
		logger.Error(err.Error())
	}
	queries := db.New(database)

	app := server.AppConfig{
		Db:       database,
		Queries:  queries,
		Logger:   logger,
		ErrorLog: errorLog,
	}
	server.StartServer(app)
}

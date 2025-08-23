package main

import (
	"github.com/mtanng9/love-my-waifu/internal/db"
	"github.com/mtanng9/love-my-waifu/internal/server"
)

func main() {
	database := db.DB{}
	db.ConnectDB(&database)
	server.StartServer()
}

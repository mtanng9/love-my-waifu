package main

import (
	"github.com/mtanng9/love-my-waifu/db"
	"github.com/mtanng9/love-my-waifu/server"
)

func main() {
	database := db.DB{}
	db.ConnectDB(&database)
	server.StartServer()
}

package main

import "github.com/mtanng9/love-my-waifu/db"

func main() {
	database := db.DB{}
	db.ConnectDB(&database)
	db.MigrateDB(&database)
}

package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mtanng9/love-my-waifu/templates/layouts"
	"github.com/mtanng9/love-my-waifu/templates/pages"
)

const PORT = 8080

func StartServer() {
	//allocate and start instantiate a new server mux
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleIndex)
	mux.HandleFunc("GET /homepage", handleUserHomePage)
	mux.HandleFunc("GET /chatpage", handleChatPage)
	mux.HandleFunc("GET /settings", handleSettingsPage)
	mux.HandleFunc("GET /favorites", handleFavoritePage)
	mux.HandleFunc("GET /chatlog", handleChatLog)

	fmt.Printf("Starting Server on PORT: %d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	page := pages.Index()
	layout := layouts.Main("Love My Waifu", page)
	layout.Render(context.TODO(), w)
}

func handleUserHomePage(w http.ResponseWriter, r *http.Request) {
	page := pages.UserHomePage()
	layout := layouts.Main("Homepage", page)
	layout.Render(context.TODO(), w)
}

func handleChatPage(w http.ResponseWriter, r *http.Request) {
	page := pages.ChatPages()
	layout := layouts.Main("Chat Page", page)
	layout.Render(context.TODO(), w)
}

func handleSettingsPage(w http.ResponseWriter, r *http.Request) {
	page := pages.SettingsPage()
	layout := layouts.Main("Settings Page", page)
	layout.Render(context.TODO(), w)
}

func handleFavoritePage(w http.ResponseWriter, r *http.Request) {
	page := pages.FavoritesPage()
	layout := layouts.Main("Favorites", page)
	layout.Render(context.TODO(), w)
}

func handleChatLog(w http.ResponseWriter, r *http.Request) {
	page := pages.ChatLog()
	layout := layouts.Main("Chat Log", page)
	layout.Render(context.TODO(), w)
}

package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mtanng9/love-my-waifu/internal/db"
	"github.com/mtanng9/love-my-waifu/internal/templates/layouts"
	"github.com/mtanng9/love-my-waifu/internal/templates/pages"
)

const PORT = 8080

type AppConfig struct {
	Db       *pgx.Conn
	Queries  *db.Queries
	Logger   *slog.Logger
	ErrorLog *log.Logger
}

func StartServer(app AppConfig) {
	//allocate and start instantiate a new server mux
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleIndex)
	mux.HandleFunc("GET /homepage", handleUserHomePage)
	mux.HandleFunc("GET /chatpage", handleChatPage)
	mux.HandleFunc("GET /settings", handleSettingsPage)
	mux.HandleFunc("GET /favorites", handleFavoritePage)
	mux.HandleFunc("GET /chatlog", handleChatLog)
	mux.HandleFunc("GET /about", handleAbout)
	mux.HandleFunc("GET /login", handleLogIn)
	mux.HandleFunc("GET /signup", handleSignUp)

	app.Logger.Info("server started", "addr", PORT)
	srv := http.Server{
		Addr:     fmt.Sprintf(":%d", PORT),
		Handler:  app.logRequest(mux),
		ErrorLog: app.ErrorLog,
	}
	err := srv.ListenAndServe()
	app.Logger.Error(err.Error())
	os.Exit(1)
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

func handleAbout(w http.ResponseWriter, r *http.Request) {
	page := pages.About()
	layout := layouts.Main("About", page)
	layout.Render(context.TODO(), w)
}

func handleLogIn(w http.ResponseWriter, r *http.Request) {
	page := pages.LogIn()
	layout := layouts.Main("Log In", page)
	layout.Render(context.TODO(), w)
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	page := pages.SignUp()
	layout := layouts.Main("Sign Up", page)
	layout.Render(context.TODO(), w)
}

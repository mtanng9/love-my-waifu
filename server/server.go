package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mtanng9/love-my-waifu/templates"
)

const PORT = 8080

func StartServer() {
	//allocate and start instantiate a new server mux
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleIndex)

	fmt.Printf("Starting Server on PORT: %d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	component := templates.Hello("Jose")
	component.Render(context.TODO(), w)
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Francky999/app-go/internal/web"
	"github.com/joho/godotenv"
)

func init() {
	// Charger .env seulement si présent (local)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Environnement de production")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// Servir les fichiers statiques (CSS, JS, images)
	assetsFS := http.FileServer(http.Dir("./views/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assetsFS))

	// Servir toutes les pages HTML depuis /views/
	http.Handle("/", http.FileServer(http.Dir("./views")))
	http.HandleFunc("/contact", web.ContactHandler)

	log.Printf("Listening on port %s\n...Click here http://localhost:%s", port, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println(err)
	}
}

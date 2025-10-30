package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Francky999/app-go/internal/web"
	"github.com/joho/godotenv"
)

func init() {
	// Charger le fichier .env seulement en développement local
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/", fs)
	http.HandleFunc("/contact", web.ContactHandler)
	log.Printf("Listening on port %s\n...Click here http://localhost:%s", port, port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

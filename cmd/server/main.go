package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Francky999/app-go/internal/web"
	"github.com/joho/godotenv"
)

// Middleware pour ajouter les headers corrects
func addHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Headers de cache et MIME types
		w.Header().Set("Cache-Control", "public, max-age=31536000") // 1 an

		// MIME types spécifiques
		if r.URL.Path[len(r.URL.Path)-4:] == ".css" {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		} else if r.URL.Path[len(r.URL.Path)-3:] == ".js" {
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		}

		h.ServeHTTP(w, r)
	})
}

func init() {
	// Charger .env seulement si présent (local)
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️ Impossible de charger .env, continuer avec les variables d'environnement")
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// Servir les fichiers statiques avec headers corrects
	fs := http.FileServer(http.Dir("./views/assets"))
	http.Handle("/assets/", addHeaders(http.StripPrefix("/assets/", fs)))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./views/index.html")
	})
	http.HandleFunc("/contact", web.ContactHandler)
	log.Printf("Listening on port %s\n...Click here http://localhost:%s", port, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

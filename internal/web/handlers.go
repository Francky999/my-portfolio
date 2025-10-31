package web

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Charger .env seulement si présent (local)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Environnement de production")
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	body := fmt.Sprintf("Nom: %s\nEmail: %s\n\nMessage:\n%s", name, email, message)

	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), os.Getenv("HOST"))

	err := smtp.SendMail(os.Getenv("HOST_PORT"), auth, os.Getenv("EMAIL"),
		[]string{os.Getenv("EMAIL")}, []byte("Subject: "+subject+"\r\n\r\n"+body))
	if err != nil {
		http.Error(w, "Erreur lors de l'envoi du mail: "+err.Error(), 500)
		return
	}
	fmt.Fprint(w, "OK")
}

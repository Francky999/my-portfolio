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
	// Headers CORS pour la production
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		log.Printf("❌ Méthode non autorisée: %s", r.Method)
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Validation des variables d'environnement
	smtpEmail := os.Getenv("EMAIL")
	smtpPassword := os.Getenv("PASSWORD")
	smtpHost := os.Getenv("HOST")
	smtpHostPort := os.Getenv("HOST_PORT")

	if smtpEmail == "" || smtpPassword == "" || smtpHost == "" || smtpHostPort == "" {
		log.Printf("❌ Variables SMTP manquantes - EMAIL: %v, PASSWORD: %v, HOST: %v, HOST_PORT: %v",
			smtpEmail != "", smtpPassword != "", smtpHost != "", smtpHostPort != "")
		http.Error(w, "Configuration email manquante", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	log.Printf("📧 Tentative d'envoi email - De: %s, Sujet: %s", email, subject)

	body := fmt.Sprintf("Nom: %s\nEmail: %s\n\nMessage:\n%s", name, email, message)

	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpHost)

	err := smtp.SendMail(smtpHostPort, auth, smtpEmail,
		[]string{smtpEmail}, []byte("Subject: "+subject+"\r\n\r\n"+body))
	if err != nil {
		log.Printf("❌ Erreur SMTP: %v", err)
		http.Error(w, "Erreur lors de l'envoi du mail", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Email envoyé avec succès")
	fmt.Fprint(w, "OK")
}

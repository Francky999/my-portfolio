package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Charger .env seulement si présent (local)
	if err := godotenv.Load(); err != nil {
		log.Println("🛠️ Environnement de production")
	}
}

// Structure pour l'API SendGrid
type SendGridEmail struct {
	Personalizations []Personalization `json:"personalizations"`
	From             EmailAddress      `json:"from"`
	Subject          string            `json:"subject"`
	Content          []Content         `json:"content"`
}

type Personalization struct {
	To []EmailAddress `json:"to"`
}

type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
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

	// Validation des variables d'environnement SendGrid
	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	fromEmail := os.Getenv("FROM_EMAIL") // Votre email vérifié dans SendGrid

	if sendgridAPIKey == "" || fromEmail == "" {
		log.Printf("❌ Variables SendGrid manquantes - API_KEY: %v, FROM_EMAIL: %v",
			sendgridAPIKey != "", fromEmail != "")
		http.Error(w, "Configuration email manquante", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	log.Printf("📧 Tentative d'envoi email via SendGrid - De: %s, Sujet: %s", email, subject)

	// Construire le message pour SendGrid
	emailBody := fmt.Sprintf("Nouveau message de contact:\n\nNom: %s\nEmail: %s\n\nMessage:\n%s", name, email, message)

	// Structure de l'email SendGrid
	sgEmail := SendGridEmail{
		Personalizations: []Personalization{
			{
				To: []EmailAddress{
					{Email: fromEmail, Name: "François OREGA"},
				},
			},
		},
		From: EmailAddress{
			Email: fromEmail,
			Name:  "Portfolio Contact",
		},
		Subject: fmt.Sprintf("Contact Portfolio: %s", subject),
		Content: []Content{
			{
				Type:  "text/plain",
				Value: emailBody,
			},
		},
	}

	// Convertir en JSON
	jsonData, err := json.Marshal(sgEmail)
	if err != nil {
		log.Printf("❌ Erreur JSON: %v", err)
		http.Error(w, "Erreur de formatage", http.StatusInternalServerError)
		return
	}

	// Envoyer via l'API SendGrid
	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("❌ Erreur création requête: %v", err)
		http.Error(w, "Erreur de requête", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+sendgridAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("❌ Erreur SendGrid API: %v", err)
		http.Error(w, "Erreur lors de l'envoi du mail", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("✅ Email envoyé avec succès via SendGrid (Status: %d)", resp.StatusCode)
		fmt.Fprint(w, "OK")
	} else {
		log.Printf("❌ Erreur SendGrid: Status %d", resp.StatusCode)
		http.Error(w, "Erreur lors de l'envoi du mail", http.StatusInternalServerError)
	}
}

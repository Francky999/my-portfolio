# Portfolio Personnel - François OREGA

Une application web moderne développée en Go pour présenter mon portfolio professionnel avec un système de contact intégré.

## 🚀 Fonctionnalités

- **Portfolio responsive** : Interface moderne et adaptative
- **Formulaire de contact** : Envoi d'emails automatisé via SMTP
- **Serveur Go** : Backend léger et performant
- **Design moderne** : Utilise Bootstrap et des animations CSS

## 🛠️ Technologies utilisées

- **Backend** : Go 1.24.4
- **Frontend** : HTML5, CSS3, JavaScript
- **Framework CSS** : Bootstrap 5
- **Serveur web** : net/http (Go standard library)
- **Email** : SMTP avec Gmail

## 📋 Prérequis

- Go 1.24.4 ou version supérieure
- Connexion internet pour les CDN et l'envoi d'emails

## 🔧 Installation

1. **Cloner le repository**

   ```bash
   git clone https://github.com/Francky999/app-go.git
   cd app-go
   ```

2. **Installer les dépendances Go**

   ```bash
   go mod tidy
   ```

3. **Configuration email** (optionnel)
   - Modifier les paramètres SMTP dans `cmd/server/main.go`
   - Remplacer les credentials Gmail par les vôtres

## 🚀 Lancement

```bash
go run cmd/server/main.go
```

L'application sera accessible sur : http://localhost:3000

## 📁 Structure du projet

```
app-go/
├── cmd/
│   └── server/
│       └── main.go          # Point d'entrée de l'application
├── views/
│   ├── assets/              # Ressources statiques (CSS, JS, images)
│   ├── index.html           # Page principale du portfolio
│   ├── portfolio-details.html
│   ├── service-details.html
│   └── starter-page.html
├── go.mod                   # Dépendances Go
├── go.sum                   # Checksums des dépendances
└── README.md               # Documentation
```

## 🌐 Pages disponibles

- **/** : Page d'accueil du portfolio
- **/contact** : Endpoint POST pour l'envoi d'emails
- **/portfolio-details.html** : Détails des projets
- **/service-details.html** : Détails des services
- **/starter-page.html** : Page de démarrage

## 📧 Fonctionnalité Contact

Le formulaire de contact permet aux visiteurs d'envoyer des messages directement via email. Les données sont traitées par l'endpoint `/contact` qui utilise le protocole SMTP pour l'envoi.

### Champs du formulaire :

- Nom
- Email
- Sujet
- Message

## 🔒 Sécurité

⚠️ **Important** : Les credentials SMTP sont actuellement en dur dans le code. Pour un environnement de production, utilisez des variables d'environnement :

```go
// Exemple d'amélioration
email := os.Getenv("SMTP_EMAIL")
password := os.Getenv("SMTP_PASSWORD")
```

## 🚀 Déploiement

Pour déployer en production :

1. **Build de l'application**

   ```bash
   go build -o portfolio cmd/server/main.go
   ```

2. **Lancement**
   ```bash
   ./portfolio
   ```

## 🤝 Contribution

Les contributions sont les bienvenues ! N'hésitez pas à :

- Ouvrir des issues pour signaler des bugs
- Proposer des améliorations
- Soumettre des pull requests

## 📄 Licence

Ce projet est sous licence MIT. Voir le fichier `LICENSE` pour plus de détails.

## 👨‍💻 Auteur

**François OREGA**

- Email : francoisorega@gmail.com
- GitHub : [@Francky999](https://github.com/Francky999)

---

⭐ N'hésitez pas à mettre une étoile si ce projet vous plaît !

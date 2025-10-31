# Portfolio Personnel - François OREGA

Une application web moderne développée en Go pour présenter mon portfolio professionnel avec un système de contact intégré utilisant SendGrid.

## 🚀 Fonctionnalités

- **Portfolio responsive** : Interface moderne et adaptative avec animations
- **Formulaire de contact** : Envoi d'emails via API SendGrid (100 emails/jour gratuits)
- **Serveur Go optimisé** : Backend léger avec gestion des assets statiques
- **Design moderne** : Bootstrap 5, AOS animations, GLightbox, Swiper
- **Déploiement cloud** : Configuration Render.com incluse
- **Logging avancé** : Suivi détaillé des opérations et erreurs

## 🛠️ Technologies utilisées

- **Backend** : Go 1.24.4
- **Frontend** : HTML5, CSS3, JavaScript (ES6+)
- **Framework CSS** : Bootstrap 5.3.3
- **Animations** : AOS (Animate On Scroll)
- **Lightbox** : GLightbox pour les images
- **Sliders** : Swiper.js
- **Email** : SendGrid API v3
- **Déploiement** : Render.com
- **Serveur web** : net/http (Go standard library)

## 📋 Prérequis

- Go 1.24.4 ou version supérieure
- Compte SendGrid (gratuit - 100 emails/jour)
- Connexion internet pour les CDN

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

3. **Configuration SendGrid**
   - Créer un compte sur [SendGrid](https://sendgrid.com)
   - Vérifier votre email dans Settings > Sender Authentication
   - Créer une API Key dans Settings > API Keys (permissions Mail Send)
   - Copier le fichier `.env.example` vers `.env`
   - Remplir les variables d'environnement :
     ```env
     PORT=3000
     SENDGRID_API_KEY=votre_cle_api_sendgrid
     FROM_EMAIL=votre_email_verifie@domain.com
     ```

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
│       └── main.go          # Point d'entrée avec serveur HTTP optimisé
├── internal/
│   └── web/
│       └── handlers.go      # Handlers HTTP (contact SendGrid)
├── views/
│   ├── assets/              # Ressources statiques servies via /assets/
│   │   ├── css/            # Styles CSS personnalisés
│   │   ├── js/             # JavaScript principal
│   │   ├── img/            # Images et favicons
│   │   └── vendor/         # Librairies tierces (Bootstrap, AOS, etc.)
│   ├── index.html          # Page principale du portfolio
│   ├── portfolio-details.html
│   ├── service-details.html
│   └── starter-page.html
├── .env                     # Variables d'environnement (local)
├── .gitignore              # Fichiers à ignorer par Git
├── render.yaml             # Configuration déploiement Render
├── go.mod                  # Dépendances Go
├── go.sum                  # Checksums des dépendances
└── README.md              # Documentation
```

## 🌐 Routes disponibles

- **GET /** : Page d'accueil du portfolio (sert `views/index.html`)
- **POST /contact** : API d'envoi d'emails via SendGrid
- **GET /assets/\*** : Fichiers statiques (CSS, JS, images)

## 📧 Système de Contact

### Fonctionnalités

- **API SendGrid** : Envoi fiable via HTTPS (pas de blocage SMTP)
- **Validation** : Vérification des champs obligatoires
- **Logging** : Suivi détaillé des envois et erreurs
- **CORS** : Support des requêtes cross-origin
- **Sécurité** : Variables d'environnement pour les clés API

### Champs du formulaire

- **Nom** : Nom du contact
- **Email** : Adresse email du contact
- **Sujet** : Objet du message
- **Message** : Contenu du message

### Format de l'email reçu

```
Nouveau message de contact:

Nom: [Nom du contact]
Email: [Email du contact]

Message:
[Contenu du message]
```

## 🔒 Sécurité

✅ **Bonnes pratiques implémentées :**

- Variables d'environnement pour les clés API
- Fichier `.env` exclu du versioning Git
- Headers CORS configurés
- Validation des données d'entrée
- Logging des erreurs sans exposer les secrets

## 🚀 Déploiement

### Déploiement local

```bash
go run cmd/server/main.go
```

### Déploiement sur Render.com

1. **Prérequis**

   - Compte Render.com (gratuit)
   - Repository GitHub connecté

2. **Configuration automatique**

   - Le fichier `render.yaml` configure automatiquement le déploiement
   - Build : `go build -o portfolio cmd/server/main.go`
   - Start : `./portfolio`

3. **Variables d'environnement à configurer dans Render**

   ```
   SENDGRID_API_KEY=votre_cle_api_sendgrid
   FROM_EMAIL=votre_email_verifie@domain.com
   ```

4. **Déploiement**
   - Connecter le repository à Render
   - Les variables sont automatiquement détectées via `render.yaml`
   - Déploiement automatique à chaque push sur `main`

### Build manuel

```bash
# Build de production
go build -o portfolio cmd/server/main.go

# Lancement
./portfolio
```

## 🔧 Configuration avancée

### Variables d'environnement

| Variable           | Description              | Requis | Défaut |
| ------------------ | ------------------------ | ------ | ------ |
| `PORT`             | Port du serveur HTTP     | Non    | `3000` |
| `SENDGRID_API_KEY` | Clé API SendGrid         | Oui    | -      |
| `FROM_EMAIL`       | Email expéditeur vérifié | Oui    | -      |

### Fichiers de configuration

- **`.env`** : Variables locales (développement)
- **`render.yaml`** : Configuration Render.com
- **`.gitignore`** : Exclusions Git (inclut `.env`)

## 🐛 Dépannage

### Problèmes courants

**Formulaire de contact ne fonctionne pas :**

- Vérifier que `SENDGRID_API_KEY` est configurée
- Vérifier que `FROM_EMAIL` est vérifié dans SendGrid
- Consulter les logs pour les erreurs détaillées

**Assets CSS/JS ne se chargent pas :**

- Vérifier que les fichiers existent dans `views/assets/`
- Vérifier les chemins dans les fichiers HTML (`/assets/...`)
- Consulter la console navigateur pour les erreurs 404

**Erreur de build :**

```bash
go mod tidy  # Nettoyer les dépendances
go build -v cmd/server/main.go  # Build verbose
```

### Logs utiles

```bash
# Logs détaillés en local
go run cmd/server/main.go

# Sur Render, consulter les logs dans le dashboard
```

## 📊 Fonctionnalités techniques

### Performance

- **Serveur HTTP natif Go** : Haute performance
- **Assets statiques optimisés** : Cache headers configurés
- **Compression** : Gzip automatique par Render

### Monitoring

- **Logging structuré** : Tous les événements tracés
- **Health check** : Route `/` pour monitoring Render
- **Error handling** : Gestion gracieuse des erreurs

### Sécurité

- **HTTPS forcé** : En production via Render
- **Variables sécurisées** : Clés API protégées
- **CORS configuré** : Headers appropriés

## 🤝 Contribution

Les contributions sont les bienvenues ! Pour contribuer :

1. **Fork** le projet
2. **Créer une branche** : `git checkout -b feature/nouvelle-fonctionnalite`
3. **Commiter** : `git commit -m 'Ajout nouvelle fonctionnalité'`
4. **Push** : `git push origin feature/nouvelle-fonctionnalite`
5. **Pull Request** : Ouvrir une PR avec description détaillée

### Standards de code

- Code Go formaté avec `gofmt`
- Tests unitaires pour les nouvelles fonctionnalités
- Documentation des fonctions publiques
- Respect des conventions Go

## 📄 Licence

Ce projet est sous licence MIT. Voir le fichier `LICENSE` pour plus de détails.

## 👨‍💻 Auteur

**François OREGA**

- 🌐 Portfolio : [hrtalenthub.blog](https://hrtalenthub.blog)
- 📧 Email : contact@hrtalenthub.blog
- 💼 LinkedIn : [François OREGA](https://www.linkedin.com/in/françois-orega-38602a11a)
- 🐙 GitHub : [@Francky999](https://github.com/Francky999)
- 🐦 Twitter : [@OregaSahore](https://x.com/OregaSahore)

## 🙏 Remerciements

- **BootstrapMade** : Template iPortfolio
- **SendGrid** : Service d'email fiable
- **Render.com** : Plateforme de déploiement
- **Go Team** : Langage performant et simple

---

⭐ **N'hésitez pas à mettre une étoile si ce projet vous plaît !**

🚀 **Déployé avec succès sur Render.com**

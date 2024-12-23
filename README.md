# README


project/
├── cmd/                # Contient les points d'entrée (ex: main.go)
│   └── main.go         # Le fichier principal de démarrage
├── graph/              # Contient les fichiers GraphQL (ex: schema.graphql)
│   └── model/          # Le fichier GraphQL pour les modèles
├── internal/           # Code spécifique à l'application (non exporté)
│   ├── database/       # Gestion de la base de données
│   ├── gql/            # GraphQL résolveurs et schémas
│   ├── models/         # Structures et types de données
│   └── utils/          # Fonctions utilitaires
├── pkg/                # Code exportable/reutilisable dans d'autres projets
├── tmp/                # Dossiers temporaires
|── .dockerignore       # ignorefile pour Docker
|── .env                # variable d'environnement
|── .air.toml           # Configuration pour Air
|── dockerfile          # Configuration pour Docker
├── go.mod              # Déclaration des dépendances
├── go.sum              # Versions exactes des dépendances
├── gqlgen              # Configuration pour gqlgen
├── Makefile            # Configuration pour Makefile
└── README
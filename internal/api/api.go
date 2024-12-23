package api

import (
	"database/sql"
	"log"
	"net/http"
	"project/conf"
	"project/graph"
	gen "project/graph"
	"project/internal/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// API encapsule la configuration, la base de données et le serveur HTTP
type API struct {
	Config *conf.Conf   // Configuration de l'application
	DB     *sql.DB      // Base de données
	Server *http.Server // Serveur HTTP
}

// NewAPI crée une nouvelle instance de l'objet API
func NewAPI(config *conf.Conf) *API {
	// Initialiser la base de données
	database := db.NewDatabase(config)

	// Créer une instance de l'objet API
	return &API{
		Config: config,
		DB:     database.DB,
		Server: &http.Server{
			Addr: ":" + config.AppPort,
		},
	}
}

// TestDB permet de tester la connexion à la base de données
func (a *API) TestDB() error {
	if err := a.DB.Ping(); err != nil {
		return err
	}

	return nil
}

// Start démarre le serveur HTTP
func (a *API) Start() error {
	log.Printf("Le serveur démarre sur http://localhost%s\n", a.Server.Addr)

	// Instanciez votre résolveur principal
	resolver := &graph.Resolver{DB: a.DB}

	// Configurez le gestionnaire GraphQL
	graphQLHandler := handler.New(gen.NewExecutableSchema(gen.Config{
		Resolvers: resolver,
	}))

	// Configurez GraphQL Playground
	playgroundHandler := playground.Handler("GraphQL Playground", "/graphql")

	// Configurez les routes HTTP
	http.Handle("/", playgroundHandler)
	http.Handle("/graphql", graphQLHandler)

	// Lancez le serveur
	return a.Server.ListenAndServe()
}

// Shutdown ferme proprement les ressources (base de données, serveur, etc.)
func (a *API) Shutdown() {
	log.Println("Fermeture de l'application...")
	if err := a.DB.Close(); err != nil {
		log.Printf("Erreur lors de la fermeture de la base de données : %v", err)
	}
}

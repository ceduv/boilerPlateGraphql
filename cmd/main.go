package main

import (
	"log"
	"net/http"
	"time"

	"project/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
)

func main() {
	// Créez le serveur GraphQL
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Ajoutez les transports nécessaires
	srv.AddTransport(transport.POST{}) // Gère les requêtes HTTP POST
	srv.AddTransport(transport.GET{})  // Gère les requêtes HTTP GET
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Permet toutes les origines (désactiver en production)
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})

	// Configurez les routes HTTP
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	// Démarrez le serveur
	log.Println("Server is running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

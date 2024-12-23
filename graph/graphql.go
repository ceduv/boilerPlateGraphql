package graph

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// NewGraphQLHandler retourne un gestionnaire HTTP pour GraphQL
func NewGraphQLHandler() http.Handler {
	// Définir un schéma GraphQL simple
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello, World!", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Erreur lors de la création du schéma GraphQL : %v", err)
	}

	// Configurer le gestionnaire GraphQL
	return handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
}

package graph

import (
	"context"
	"database/sql"
	"fmt"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver regroupe les dépendances nécessaires pour exécuter les résolveurs
type Resolver struct {
	DB *sql.DB // Connexion à la base de données
}

// PingDB est un exemple de résolveur pour tester la connexion à la base de données
func (r *Resolver) PingDB(ctx context.Context) (string, error) {
	if err := r.DB.Ping(); err != nil {
		return "", fmt.Errorf("Erreur lors du test de la base de données : %v", err)
	}

	return "La connexion à la base de données est opérationnelle !", nil
}

// Test est un exemple de résolveur simple
func (r *Resolver) Test(ctx context.Context) (string, error) {
	return "Le serveur GraphQL fonctionne correctement !", nil
}

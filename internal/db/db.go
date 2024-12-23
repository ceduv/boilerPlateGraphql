package db

import (
	"database/sql"
	"fmt"
	"log"

	conf "project/conf"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

type DB struct {
	DB *sql.DB
}

// NewDatabase initialise et retourne une connexion à la base de données
func NewDatabase(cfg *conf.Conf) *DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("[DB] error open: %v", err)
	}

	// Vérifier la connexion
	err = db.Ping()
	if err != nil {
		log.Fatalf("[DB] ping fail %v", err)
	}

	log.Println("[DB] OK")
	return &DB{DB: db}
}

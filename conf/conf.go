package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Conf struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

// LoadConf charge la configuration à partir du fichier .env
func LoadConf() *Conf {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}

	return &Conf{
		AppPort:    os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSL_MODE"),
	}
}

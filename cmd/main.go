package main

import (
	"log"
	"os"
	"os/signal"
	"project/conf"
	"project/internal/api"
	"syscall"
)

func main() {
	config := conf.LoadConf()
	app := api.NewAPI(config)

	// Tester la connexion à la base de données
	if err := app.TestDB(); err != nil {
		log.Fatalf("Erreur de connexion à la base de données : %v", err)
	}

	// Lancer le serveur HTTP
	go func() {
		if err := app.Start(); err != nil {
			log.Fatalf("Erreur lors du lancement du serveur : %v", err)
		}
	}()

	// Attendre les interruptions pour arrêter proprement
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	app.Shutdown()
	log.Println("Application arrêtée.")
}

// migrate.go

package main

import (
	"ProjetGo/db"
	"ProjetGo/models"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la base de données: %v", err)
	}

	// Crée ou met à jour la table en fonction du modèle User.
	if err := database.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Erreur lors de la création ou mise à jour de la table: %v", err)
	}

	log.Println("Migration terminée avec succès!")
}

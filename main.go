package main

import (
	"log"
	"net/http"
	"os"

	"ProjetGo/controllers"
	"ProjetGo/db"
	"github.com/gorilla/mux"
)

func main() {
	db.LoadEnv() // Charger les variables d'environnement

	// Initialisation de la base de données
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la base de données: %v", err)
	}
	defer func() {
		sqlDB, _ := database.DB()
		sqlDB.Close()
	}()

	// Initialise la variable DB globale dans le package controllers
	controllers.DB = database

	// Configuration du router
	r := mux.NewRouter()

	// Routes pour les utilisateurs
	r.HandleFunc("/users/{id:[0-9]+}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")

	// Utilisez le middleware d'authentification
	// r.Handle("/users/{id:[0-9]+}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetUser))).Methods("GET")

	// Démarrage du serveur
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Le serveur est démarré sur le port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

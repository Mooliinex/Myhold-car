package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// LoadEnv charge les variables d'environnement depuis le fichier .env
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier .env: %v", err)
	}
}

// InitDB initialise et retourne une connexion à la base de données PostgreSQL
func InitDB() (*gorm.DB, error) {
	LoadEnv() // Chargez les variables d'environnement depuis le fichier .env

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateDatabase crée une nouvelle base de données PostgreSQL
func CreateDatabase(databaseName string) error {
	LoadEnv() // Chargez les variables d'environnement depuis le fichier .env

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PORT"),
	)

	// Ouvrir une connexion à PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Assurez-vous de fermer la connexion à la fin
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Exécutez la requête SQL pour créer la base de données
	if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", databaseName)).Error; err != nil {
		return err
	}

	return nil
}

package main

import (
	"log"

	"github.com/Lycee_Connect/config"
	"github.com/Lycee_Connect/models"
	"github.com/Lycee_Connect/router"
)

func main() {
	// Connexion DB
	config.Connect()
	//  Migration automatique
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}
	//   Router
	r := router.Setup()
	log.Println("HTTP server on :8080")
	r.Run(":8080")
}

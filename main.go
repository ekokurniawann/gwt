package main

import (
	"log"
	"net/http"

	"go-web-template/server"
	"go-web-template/server/configs"
	"go-web-template/server/database"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	db, err := database.NewConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
	}

	router := routes.NewRouter(db)

	port := cfg.Server.Port

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

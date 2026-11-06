package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"transaction-api/config"
	"transaction-api/handlers"
	"transaction-api/repositories"
	"transaction-api/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	repository := repositories.NewTransactionRepository(db)
	service := services.NewTransactionService(repository)
	handler := handlers.NewTransactionHandler(service)

	http.HandleFunc("/transactions", handler.HandleCollection)
	http.HandleFunc("/transactions/", handler.HandleItem)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

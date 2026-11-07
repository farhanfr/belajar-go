package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"transaction-api/config"
	"transaction-api/handlers"
	"transaction-api/middlewares"
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

	mux := http.NewServeMux()

	mux.HandleFunc("/transactions", handler.HandleCollection)
	mux.HandleFunc("/transactions/", handler.HandleItem)
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("testing recovery middleware")
	})

	app := middlewares.Logging(
		middlewares.Recovery(mux),
	)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           app,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Server running on :8080")
	log.Fatal(server.ListenAndServe())
}

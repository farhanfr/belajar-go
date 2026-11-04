package main

import (
	"fmt"
	"net/http"

	"transaction-api/handlers"
	"transaction-api/repositories"
	"transaction-api/services"
)

func main() {
	repo := repositories.NewTransactionRepository()
	service := services.NewTransactionService(repo)
	handler := handlers.NewTransactionHandler(service)

	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetAll(w, r)
		case http.MethodPost:
			handler.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/transactions/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetByID(w, r)
			return
		} else if r.Method == http.MethodPut {
			handler.Update(w, r)
			return
		} else if r.Method == http.MethodDelete {
			handler.Delete(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)
}

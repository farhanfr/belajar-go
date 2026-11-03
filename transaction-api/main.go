package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Transaction struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Amount int    `json:"amount"`
}

var transactions = []Transaction{
	{ID: 1, Title: "Beli Baju", Amount: 100000},
	{ID: 2, Title: "Beli Celana", Amount: 200000},
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func getTransactionsByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/transactions/"))

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	transaction, found := findTransaction(transactions, id)

	if !found {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	transaction.ID = len(transactions) + 1
	transactions = append(transactions, transaction)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func updateTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/transactions/"))

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var request Transaction

	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, transaction := range transactions {
		if transaction.ID == id {
			transactions[i].Title = request.Title
			transactions[i].Amount = request.Amount

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(transactions[i])
			return
		}
	}

	http.Error(w, "Transaction not found", http.StatusNotFound)
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/transactions/"))

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, transaction := range transactions {
		if transaction.ID == id {
			transactions = append(
				transactions[:i],
				transactions[i+1:]...,
			)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Transaction not found", http.StatusNotFound)
}

func findTransaction(transactions []Transaction, id int) (Transaction, bool) {
	for _, transaction := range transactions {
		if transaction.ID == id {
			return transaction, true
		}
	}

	return Transaction{}, false
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	isCollection := r.URL.Path == "/transactions"
	isDetail := strings.HasPrefix(r.URL.Path, "/transactions/")

	switch r.Method {
	case http.MethodGet:
		if isCollection {
			getTransactions(w, r)
		} else if isDetail {
			getTransactionsByID(w, r)
		} else {
			http.NotFound(w, r)
		}

	case http.MethodPost:
		if isCollection {
			createTransaction(w, r)
		} else {
			http.NotFound(w, r)
		}

	case http.MethodPut:
		if isDetail {
			updateTransaction(w, r)
		} else {
			http.NotFound(w, r)
		}

	case http.MethodDelete:
		if isDetail {
			deleteTransaction(w, r)
		} else {
			http.NotFound(w, r)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/transactions", transactionHandler)
	http.HandleFunc("/transactions/", transactionHandler)

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

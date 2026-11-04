package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"transaction-api/config"
	"transaction-api/models"
)

func getTransactions(db *sql.DB) {
	rows, err := db.Query(
		"SELECT id, title, amount FROM transactions ORDER BY id",
	)

	if err != nil {
		log.Fatal("Query failed:", err)
	}

	defer rows.Close()

	var transactions []models.Transaction

	for rows.Next() {
		var transaction models.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.Title,
			&transaction.Amount,
		)

		if err != nil {
			log.Fatal("Scan failed:", err)
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Rows error:", err)
	}

	for _, transaction := range transactions {
		fmt.Println(
			"ID:", transaction.ID,
			"Title:", transaction.Title,
			"Amount:", transaction.Amount,
		)
	}
}

func main() {
	if err := godotenv.Overload(); err != nil {
		log.Fatal(err)
	}

	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer db.Close()

	fmt.Println("PostgreSQL connected successfully!")

	getTransactions(db)
}

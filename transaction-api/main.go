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

func getTransactionByID(db *sql.DB, id int) {
	var transaction models.Transaction

	err := db.QueryRow(
		"SELECT id, title, amount FROM transactions WHERE id = $1",
		id,
	).Scan(
		&transaction.ID,
		&transaction.Title,
		&transaction.Amount,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Transaction not found")
			return
		}

		log.Println("Query failed:", err)
		return
	}

	fmt.Println("ID:", transaction.ID)
	fmt.Println("Title:", transaction.Title)
	fmt.Println("Amount:", transaction.Amount)
}

func createTransaction(db *sql.DB, transaction models.Transaction) (int, error) {
	var id int

	err := db.QueryRow(
		`INSERT INTO transactions (title, amount)
         VALUES ($1, $2)
         RETURNING id`,
		transaction.Title,
		transaction.Amount,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func updateTransaction(db *sql.DB, id int, transaction models.Transaction) error {
	result, err := db.Exec(
		`UPDATE transactions
         SET title = $1, amount = $2
         WHERE id = $3`,
		transaction.Title,
		transaction.Amount,
		id,
	)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("transaction not found")
	}

	return nil
}

func deleteTransaction(db *sql.DB, id int) error {
	result, err := db.Exec(
		"DELETE FROM transactions WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("transaction not found")
	}

	return nil
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

	getTransactionByID(db, 1)
}

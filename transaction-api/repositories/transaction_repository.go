package repositories

import (
	"context"
	"database/sql"
	"transaction-api/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Get all transactions
func (r *TransactionRepository) FindAll(ctx context.Context) ([]models.Transaction, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, title, amount FROM transactions ORDER BY id",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transactions := make([]models.Transaction, 0)

	for rows.Next() {
		var transaction models.Transaction

		if err := rows.Scan(
			&transaction.ID,
			&transaction.Title,
			&transaction.Amount,
		); err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

// Get transaction by ID
func (r *TransactionRepository) FindByID(ctx context.Context, id int) (models.Transaction, error) {
	var transaction models.Transaction

	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, title, amount FROM transactions WHERE id = $1",
		id,
	).Scan(
		&transaction.ID,
		&transaction.Title,
		&transaction.Amount,
	)

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

// Create transaction
func (r *TransactionRepository) Create(ctx context.Context, transaction models.Transaction) (models.Transaction, error) {
	err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO transactions (title, amount)
         VALUES ($1, $2)
         RETURNING id`,
		transaction.Title,
		transaction.Amount,
	).Scan(&transaction.ID)

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

// Update transaction
func (r *TransactionRepository) Update(ctx context.Context, id int, transaction models.Transaction) (models.Transaction, error) {
	err := r.db.QueryRowContext(
		ctx,
		`UPDATE transactions
         SET title = $1, amount = $2
         WHERE id = $3
         RETURNING id, title, amount`,
		transaction.Title,
		transaction.Amount,
		id,
	).Scan(
		&transaction.ID,
		&transaction.Title,
		&transaction.Amount,
	)

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

// Delete transaction
func (r *TransactionRepository) Delete(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(
		ctx,
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
		return sql.ErrNoRows
	}

	return nil
}

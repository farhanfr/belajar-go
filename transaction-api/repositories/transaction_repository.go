package repositories

import "transaction-api/models"

type TransactionRepository struct {
	transactions []models.Transaction
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		transactions: []models.Transaction{
			{ID: 1, Title: "Beli Baju", Amount: 100000},
			{ID: 2, Title: "Beli Celana", Amount: 200000},
		},
	}
}

func (r *TransactionRepository) FindAll() []models.Transaction {
	return r.transactions
}

func (r *TransactionRepository) FindByID(id int) (models.Transaction, bool) {
	for _, transaction := range r.transactions {
		if transaction.ID == id {
			return transaction, true
		}
	}

	return models.Transaction{}, false
}

func (r *TransactionRepository) Create(transaction models.Transaction) models.Transaction {
	transaction.ID = len(r.transactions) + 1
	r.transactions = append(r.transactions, transaction)

	return transaction
}

func (r *TransactionRepository) Update(id int, updated models.Transaction) (models.Transaction, bool) {
	for i, transaction := range r.transactions {
		if transaction.ID == id {
			r.transactions[i].Title = updated.Title
			r.transactions[i].Amount = updated.Amount

			return r.transactions[i], true
		}
	}

	return models.Transaction{}, false
}

func (r *TransactionRepository) Delete(id int) {
	for i, t := range r.transactions {
		if t.ID == id {
			r.transactions = append(r.transactions[:i], r.transactions[i+1:]...)
			return
		}
	}
}

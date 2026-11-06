package services

import (
	"errors"
	"strings"

	"transaction-api/models"
	"transaction-api/repositories"
)

type TransactionService struct {
	repository *repositories.TransactionRepository
}

func NewTransactionService(repository *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{
		repository: repository,
	}
}

// Get all transactions
func (s *TransactionService) GetAll() ([]models.Transaction, error) {
	return s.repository.FindAll()
}

// Get transaction by ID
func (s *TransactionService) GetByID(id int) (models.Transaction, error) {
	if id <= 0 {
		return models.Transaction{}, errors.New("invalid transaction ID")
	}

	return s.repository.FindByID(id)
}

// Create transaction
func (s *TransactionService) Create(transaction models.Transaction) (models.Transaction, error) {
	transaction.Title = strings.TrimSpace(transaction.Title)

	if transaction.Title == "" {
		return models.Transaction{}, errors.New("title is required")
	}

	if transaction.Amount <= 0 {
		return models.Transaction{}, errors.New("amount must be greater than zero")
	}

	return s.repository.Create(transaction)
}

// Update transaction
func (s *TransactionService) Update(id int, transaction models.Transaction) (models.Transaction, error) {
	if id <= 0 {
		return models.Transaction{}, errors.New("invalid transaction ID")
	}

	transaction.Title = strings.TrimSpace(transaction.Title)

	if transaction.Title == "" {
		return models.Transaction{}, errors.New("title is required")
	}

	if transaction.Amount <= 0 {
		return models.Transaction{}, errors.New("amount must be greater than zero")
	}

	return s.repository.Update(id, transaction)
}

// Delete transaction
func (s *TransactionService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid transaction ID")
	}

	return s.repository.Delete(id)
}

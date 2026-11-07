package services

import (
	"context"
	"errors"
	"strings"

	"transaction-api/models"
)

type TransactionRepository interface {
	FindAll(ctx context.Context) ([]models.Transaction, error)
	FindByID(ctx context.Context, id int) (models.Transaction, error)
	Create(ctx context.Context, transaction models.Transaction) (models.Transaction, error)
	Update(ctx context.Context, id int, transaction models.Transaction) (models.Transaction, error)
	Delete(ctx context.Context, id int) error
}

type TransactionService struct {
	repository TransactionRepository
}

func NewTransactionService(repository TransactionRepository) *TransactionService {
	return &TransactionService{
		repository: repository,
	}
}

func (s *TransactionService) GetAll(ctx context.Context) ([]models.Transaction, error) {
	return s.repository.FindAll(ctx)
}

func (s *TransactionService) GetByID(ctx context.Context, id int) (models.Transaction, error) {
	if id <= 0 {
		return models.Transaction{}, errors.New("invalid transaction ID")
	}

	return s.repository.FindByID(ctx, id)
}

func (s *TransactionService) Create(ctx context.Context, transaction models.Transaction) (models.Transaction, error) {
	transaction.Title = strings.TrimSpace(transaction.Title)

	if transaction.Title == "" {
		return models.Transaction{}, errors.New("title is required")
	}

	if transaction.Amount <= 0 {
		return models.Transaction{}, errors.New("amount must be greater than zero")
	}

	return s.repository.Create(ctx, transaction)
}

func (s *TransactionService) Update(ctx context.Context, id int, transaction models.Transaction) (models.Transaction, error) {
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

	return s.repository.Update(ctx, id, transaction)
}

func (s *TransactionService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid transaction ID")
	}

	return s.repository.Delete(ctx, id)
}

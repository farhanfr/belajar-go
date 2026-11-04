package services

import (
	"errors"
	"transaction-api/models"
	"transaction-api/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) GetAll() []models.Transaction {
	return s.repo.FindAll()
}

func (s *TransactionService) GetByID(id int) (models.Transaction, error) {
	transaction, found := s.repo.FindByID(id)

	if !found {
		return models.Transaction{}, errors.New("transaction not found")
	}

	return transaction, nil
}

func (s *TransactionService) Create(transaction models.Transaction) models.Transaction {
	return s.repo.Create(transaction)
}

func (s *TransactionService) Update(id int, transaction models.Transaction) (models.Transaction, error) {
	updated, found := s.repo.Update(id, transaction)

	if !found {
		return models.Transaction{}, errors.New("transaction not found")
	}

	return updated, nil
}

func (s *TransactionService) Delete(id int) {
	s.repo.Delete(id)
}

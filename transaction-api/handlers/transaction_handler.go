package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"transaction-api/models"
	"transaction-api/services"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

// GET /transactions
// POST /transactions
func (h *TransactionHandler) HandleCollection(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	switch r.Method {
	case http.MethodGet:
		transactions, err := h.service.GetAll(ctx)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to get transactions")
			return
		}

		writeJSON(w, http.StatusOK, transactions)

	case http.MethodPost:
		var transaction models.Transaction

		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			writeError(w, http.StatusBadRequest, "invalid request body")
			return
		}

		created, err := h.service.Create(ctx, transaction)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		writeJSON(w, http.StatusCreated, created)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// GET /transactions/{id}
// PUT /transactions/{id}
// DELETE /transactions/{id}
func (h *TransactionHandler) HandleItem(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/transactions/")
	id, err := strconv.Atoi(idString)

	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid transaction ID")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	switch r.Method {
	case http.MethodGet:
		transaction, err := h.service.GetByID(ctx, id)
		if err != nil {
			handleServiceError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, transaction)

	case http.MethodPut:
		var transaction models.Transaction

		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			writeError(w, http.StatusBadRequest, "invalid request body")
			return
		}

		updated, err := h.service.Update(ctx, id, transaction)
		if err != nil {
			handleServiceError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		if err := h.service.Delete(ctx, id); err != nil {
			handleServiceError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleServiceError(w http.ResponseWriter, err error) {
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, "transaction not found")
		return
	}

	writeError(w, http.StatusBadRequest, err.Error())
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}

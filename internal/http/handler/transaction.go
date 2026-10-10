package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	transactiondto "github.com/gorkagg10/equity-calculator-api/internal/http/dto/transaction"
	"github.com/gorkagg10/equity-calculator-api/internal/service"
)

type Transaction struct {
	service *service.Transaction
}

func NewTransaction(service *service.Transaction) *Transaction {
	return &Transaction{
		service: service,
	}
}

func (t *Transaction) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", t.Add)

	return router
}

func (t *Transaction) Add(w http.ResponseWriter, r *http.Request) {
	var requestBody transactiondto.AddTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_BODY", "malformed JSON body")
		return
	}
	requestPortfolioID := chi.URLParam(r, "portfolioID")
	portfolioID, err := uuid.Parse(requestPortfolioID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_BODY", "invalid portfolioID")
		return
	}
	assetID, err := uuid.Parse(requestBody.AssetID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_BODY", "invalid assetID")
		return
	}

	transaction, err := t.service.Add(
		r.Context(),
		portfolioID,
		assetID,
		requestBody.UnitPrice,
		requestBody.Quantity,
		requestBody.Currency,
		requestBody.TransactionType,
	)

	response := transactiondto.AddTransactionResponse{
		ID: transaction.ID().String(),
	}

	writeJSON(w, http.StatusCreated, response)
}

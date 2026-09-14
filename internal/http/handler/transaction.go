package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	transactiondto "github.com/gorkagg10/equity-calculator-api/internal/http/dto/transaction"
)

type Transaction struct{}

func NewTransaction() *Transaction {
	return &Transaction{}
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
	writeJSON(w, http.StatusCreated, nil)
}

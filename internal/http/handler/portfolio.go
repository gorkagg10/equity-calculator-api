package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	portfoliodto "github.com/gorkagg10/equity-calculator-api/internal/http/dto/portfolio"
)

type Portfolio struct{}

func NewPortfolio() *Portfolio {
	return &Portfolio{}
}

func (p *Portfolio) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/portfolios", p.Add)

	return router
}

func (p *Portfolio) Add(w http.ResponseWriter, r *http.Request) {
	var request portfoliodto.AddPortfolioRequest
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	response := portfoliodto.AddPortfolioResponse{
		ID: "hola",
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}

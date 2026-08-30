package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	portfoliodto "github.com/gorkagg10/equity-calculator-api/internal/http/dto/portfolio"
	"github.com/gorkagg10/equity-calculator-api/internal/service"
)

type Portfolio struct {
	service *service.Portfolio
}

func NewPortfolio(service *service.Portfolio) *Portfolio {
	return &Portfolio{
		service: service,
	}
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
	if err := request.Validate(); err != nil {
		return
	}
	response := portfoliodto.AddPortfolioResponse{
		ID: "hola",
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}

func writeError(w http.ResponseWriter, status int)

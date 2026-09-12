package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
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
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_BODY", "malformed JSON body")
		return
	}
	if err := request.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_BODY", err.Error())
		return
	}

	portfolio, err := p.service.AddPortfolio(request.Name)
	if err != nil {
		handleServiceError(w, err)
	}
	response := portfoliodto.AddPortfolioResponse{
		ID: portfolio.ID().String(),
	}
	writeJSON(w, http.StatusCreated, response)
}

func handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidName):
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
	default:
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "unexpected error")
	}
}

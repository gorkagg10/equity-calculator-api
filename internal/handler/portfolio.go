package handler

import "github.com/go-chi/chi/v5"

type Portfolio struct{}

func NewPortfolio() *Portfolio {
	return &Portfolio{}
}

func (p *Portfolio) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/portfolios/", nil)

	return router
}

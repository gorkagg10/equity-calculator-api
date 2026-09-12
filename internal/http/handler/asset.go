package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	assetdto "github.com/gorkagg10/equity-calculator-api/internal/http/dto/asset"
)

type Asset struct{}

func (a *Asset) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/portfolios", a.Add)

	return router
}

func (a *Asset) Add(w http.ResponseWriter, r *http.Request) {
	var requestBody assetdto.AddAssetRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_BODY", "malformed JSON body")
		return
	}
	if err := requestBody.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_BODY", err.Error())
		return
	}

	response := assetdto.AddAssetResponse{
		ID: uuid.NewString(),
	}
	writeJSON(w, http.StatusCreated, response)
}

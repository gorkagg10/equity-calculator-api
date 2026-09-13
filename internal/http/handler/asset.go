package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	assetdto "github.com/gorkagg10/equity-calculator-api/internal/http/dto/asset"
	"github.com/gorkagg10/equity-calculator-api/internal/service"
)

type Asset struct {
	service *service.Asset
}

func NewAsset(service *service.Asset) *Asset {
	return &Asset{
		service: service,
	}
}

func (a *Asset) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", a.Add)

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
	asset, err := a.service.Add(r.Context(), requestBody.Symbol)
	if err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_BODY", err.Error())
		return
	}

	response := assetdto.AddAssetResponse{
		ID: asset.ID().String(),
	}
	writeJSON(w, http.StatusCreated, response)
}

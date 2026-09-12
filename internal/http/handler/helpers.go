package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorkagg10/equity-calculator-api/pkg/apierror"
)

func writeJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		return fmt.Errorf("error writing json: %w", err)
	}
	return nil
}

func writeError(w http.ResponseWriter, status int, code, message string) error {
	return writeJSON(w, status, apierror.New(code, message))
}

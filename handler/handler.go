package handler

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"simple-rest-service/config"
)

type Handler struct {
	config *config.Config
}

type ExternalResponse struct {
	Supply []struct {
		Amount string `json:"amount"`
	} `json:"supply"`
}

type Response struct {
	Amount big.Int `json:"amount"`
}

// NewHandler creates a new Handler with the given configuration.
func NewHandler(c *config.Config) *Handler {
	h := new(Handler)
	h.config = c
	return h
}

// httpInternalError writes an HTTP internal server error to the writer and logs an error message to the standard logger.
func httpInternalError(w http.ResponseWriter, logMessage string, err error) {
	http.Error(w, logMessage, http.StatusInternalServerError)
	log.Println(logMessage, err)
}

// HandleGetTotalSupply is an HTTP request handler that fetches the total supply of NGL tokens from an external service, unmarshals the JSON response, and returns the amount in a JSON response.
func (h *Handler) HandleGetTotalSupply(w http.ResponseWriter, _ *http.Request) {
	respData, err := http.Get(h.config.URL)
	if err != nil {
		httpInternalError(w, "Failed to fetch data", err)
		return
	}
	defer respData.Body.Close()

	var externalResponse ExternalResponse
	if err = json.NewDecoder(respData.Body).Decode(&externalResponse); err != nil {
		httpInternalError(w, "Failed to unmarshal external fetch data json", err)
		return
	}

	response := &Response{
		Amount: big.Int{},
	}
	response.Amount.SetString(externalResponse.Supply[0].Amount, 10)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		httpInternalError(w, "Failed to marshal response json", err)
		return
	}
}

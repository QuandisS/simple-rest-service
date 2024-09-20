package handler

import (
	"encoding/json"
	"io"
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

func NewHandler(c *config.Config) *Handler {
	h := new(Handler)
	h.config = c
	return h
}

func httpInternalError(w http.ResponseWriter, logMessage string, err error) {
	http.Error(w, logMessage, http.StatusInternalServerError)
	log.Println(logMessage, err)
}

func (h *Handler) GetTotalSuppplyHandler(w http.ResponseWriter, _ *http.Request) {
	respData, err := http.Get(h.config.URL)
	if err != nil {
		httpInternalError(w, "Failed to fetch data", err)
		return
	}
	defer respData.Body.Close()

	body, err := io.ReadAll(respData.Body)
	if err != nil {
		httpInternalError(w, "Failed to read body", err)
		return
	}

	var externalResponse ExternalResponse
	if err = json.Unmarshal(body, &externalResponse); err != nil {
		httpInternalError(w, "Failed to unmarshal external fetch data json", err)
		return
	}

	response := &Response{
		Amount: big.Int{},
	}
	response.Amount.SetString(externalResponse.Supply[0].Amount, 10)

	jsonData, err := json.Marshal(response)
	if err != nil {
		httpInternalError(w, "Failed to marshal response json", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}

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

func (h *Handler) GetTotalSuppplyHandler(w http.ResponseWriter, _ *http.Request) {
	respData, err := http.Get(h.config.URL)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		log.Println("Error when fetching data:", err)
		return
	}
	defer respData.Body.Close()

	body, err := io.ReadAll(respData.Body)
	if err != nil {
		http.Error(w, "Error when reading body", http.StatusInternalServerError)
		log.Println("Error when reading body:", err)
		return
	}

	var externalResponse ExternalResponse
	if err = json.Unmarshal(body, &externalResponse); err != nil {
		http.Error(w, "Error when unmarshaling json", http.StatusInternalServerError)
		log.Println("Error when unmarshaling json:", err)
		return
	}

	response := &Response{
		Amount: big.Int{},
	}
	response.Amount.SetString(externalResponse.Supply[0].Amount, 10)

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error when marshaling json", http.StatusInternalServerError)
		log.Println("Error when marshaling json:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}

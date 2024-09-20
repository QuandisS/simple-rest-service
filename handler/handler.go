package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"simple-rest-service/config"
)

type Handler struct {
	config *config.Config
}

type Response struct {
	Supply []struct {
		Amount string `json:"amount"`
	} `json:"supply"`
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

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Error when unmarshaling json", http.StatusInternalServerError)
		log.Println("Error when unmarshaling json:", err)
		return
	}

	jsonData := fmt.Sprintf("{\"amount\": %s}", response.Supply[0].Amount)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}

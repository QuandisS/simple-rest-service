package handler

import (
	"encoding/json"
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

	log.Println(string(body))

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Error when unmarshaling json", http.StatusInternalServerError)
		log.Println("Error when unmarshaling json:", err)
		return
	}

	jsonData, err := json.Marshal(response.Supply[0])
	if err != nil {
		http.Error(w, "Error when marshaling response json", http.StatusInternalServerError)
		log.Println("Error when marshaling response json:", err)
		return
	}

	log.Println(string(jsonData))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

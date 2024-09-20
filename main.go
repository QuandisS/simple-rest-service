package main

import (
	"log"
	"net/http"
	"simple-rest-service/config"
	"simple-rest-service/handler"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler(config)

	http.HandleFunc("GET /getTotalSupply", h.GetTotalSuppplyHandler)

	if err = http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}

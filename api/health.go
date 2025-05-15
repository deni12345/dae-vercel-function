package api

import (
	"log"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("App's health is good")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Everything is good"))
}

package api

import (
	"log"
	"net/http"

	"github.com/dae-vercel-function/cloud"
)

// @Summary      Health check endpoint
// @Description  Check if the application is running and healthy
// @Tags         health
// @Accept       json
// @Produce      plain
// @Success      200  {string}  string  "Everything is good"
// @Router       /health [get]
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("App's health is good")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Everything is good")); err != nil {
		cloud.LogError("Failed to write response: %v", err)
	}
}

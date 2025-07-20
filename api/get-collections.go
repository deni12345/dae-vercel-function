package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dae-vercel-function/cloud"
)

// @Summary      Get Firestore collections count
// @Description  Returns the number of collections in the Firestore database
// @Tags         collections
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]int  "Collections count response"
// @Failure      500  {string}  string          "Internal Server Error"
// @Router       /collections [get]
func GetCollectionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	firestore := cloud.NewFireStore(r.Context(), "drink-and-eat-b7e64")
	defer firestore.Close()

	collections, err := firestore.GetCollections(r.Context())
	if err != nil {
		cloud.LogError("Failed to verify Firestore client initialization: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Printf("Firebase app created successfully. Collection count: %d", len(collections))
	if err := json.NewEncoder(w).Encode(map[string]int{
		"collections-count": len(collections),
	}); err != nil {
		cloud.LogError("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

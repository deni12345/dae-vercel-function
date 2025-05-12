package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dae-vercel-function/cloud"
)

func SheetSubscribeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	firestore := cloud.NewFireStore(r.Context(), "drink-and-eat-b7e64")
	defer firestore.Close()

	collections, err := firestore.GetCollections(r.Context())
	if err != nil {
		cloud.LogError("Failed to verify Firestore client initialization: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Printf("Firebase app created successfully. Collection path: %d", len(collections))
	if err := json.NewEncoder(w).Encode(map[string]int{
		"collections-count": len(collections),
	}); err != nil {
		cloud.LogError("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

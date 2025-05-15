package api

import (
	"encoding/json"
	"net/http"

	"github.com/dae-vercel-function/cloud"
)

func SheetListenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	sheetID := params.Get("sheetId")

	firestore := cloud.NewFireStore(r.Context(), "drink-and-eat-b7e64")
	defer firestore.Close()

	changes, err := firestore.ObservceCollection(r.Context(), sheetID)
	if err != nil {
		cloud.LogError("Failed to verify Firestore client initialization: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(changes); err != nil {
		cloud.LogError("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

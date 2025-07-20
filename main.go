package main

import (
	"fmt"
	"net/http"

	"github.com/dae-vercel-function/api"
)

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/api/broadcast-sheet", api.BroadcastSheetHandler)
	http.HandleFunc("/api/collections", api.GetCollectionsHandler)
	http.HandleFunc("/api/health", api.HealthHandler)
	http.HandleFunc("/api/swagger/*",api.SwaggerHandler)


	// Index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "index.html")
	})

	port := ":3000"
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

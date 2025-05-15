package main

import (
	"fmt"
	"net/http"

	"github.com/dae-vercel-function/api"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/api/sheet-listen", api.SheetListenHandler)
	http.HandleFunc("/api/sheet-subscribe", api.SheetSubscribeHandler)
	http.HandleFunc("/api/health", api.HealthHandler)

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

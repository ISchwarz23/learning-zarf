package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// GreetingResponse defines the JSON structure returned by /api/greeting
type GreetingResponse struct {
	Greeting string `json:"greeting"`
}

func main() {
	greeting := "Hello, World!"

	http.HandleFunc("/api/greeting", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GreetingResponse{Greeting: greeting})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

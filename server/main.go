package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/atanda0x/classify-number/internal/handlers"
)

const PORT = "8000"

func main() {
	http.HandleFunc("/api/classify-number", handlers.Classify)

	port := os.Getenv("PORT")
	if strings.TrimSpace(port) == "" {
		port = PORT
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

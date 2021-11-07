package main

import (
	"log"
	"net/http"

	"github.com/kramuenk/platform-test/app/internal/handlers"
)

func main() {
	log.Println("starting server on port 8080")
	http.HandleFunc("/", handlers.NewHello())
	http.HandleFunc("/status", handlers.NewStatus())
	http.HandleFunc("/metadata", handlers.NewMetadata())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

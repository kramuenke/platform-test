package main

import (
	"log"
	"net/http"

	"github.com/kramuenk/platform-test/app/internal/handlers"
)

func main() {
	log.Println("starting server on port 8080")
	http.HandleFunc("/", handlers.NewHello())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

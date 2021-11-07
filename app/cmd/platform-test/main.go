package main

import (
	"log"
	"net/http"

	"github.com/kramuenk/platform-test/app/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.NewHello())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

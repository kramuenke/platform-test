package handlers

import (
	"fmt"
	"net/http"
)

func NewHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	}
}

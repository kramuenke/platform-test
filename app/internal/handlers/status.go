package handlers

import (
	"net/http"
)

func NewStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

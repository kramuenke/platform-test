package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

type metadata struct {
	Version     string
	Description string
	CommitSha   string
}

func NewMetadata() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := new(metadata)
		data.Version = "1.0"
		data.Description = "platform test app"
		data.CommitSha = os.Getenv("COMMIT_SHA")
		json.NewEncoder(w).Encode(data)
	}
}

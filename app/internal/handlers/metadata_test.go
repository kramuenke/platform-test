package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMetadataHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	NewMetadata()(res, req)

	expectedCode := http.StatusOK
	if res.Code != expectedCode {
		t.Errorf("got status %d but expected %d", res.Code, expectedCode)
	}

	expectedMetadata := metadata{
		Version:     "1.0",
		Description: "platform test app",
		CommitSha:   "1234",
	}

	var actualMetadata metadata
	err := json.NewDecoder(res.Body).Decode(&actualMetadata)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(expectedMetadata, actualMetadata); diff != "" {
		t.Errorf("metadata mismatch: \n%s", diff)
	}
}

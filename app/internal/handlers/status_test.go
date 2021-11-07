package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	NewStatus()(res, req)

	expectedCode := http.StatusOK
	if res.Code != expectedCode {
		t.Errorf("got status %d but expected %d", res.Code, expectedCode)
	}
}

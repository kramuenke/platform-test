package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	NewHello()(res, req)

	expectedCode := http.StatusOK
	if res.Code != expectedCode {
		t.Errorf("got status %d but expected %d", res.Code, expectedCode)
	}

	expectedBody := "Hello World"
	if res.Body.String() != expectedBody {
		t.Errorf("unexpected body: got %v want %v", res.Body.String(), expectedBody)
	}
}

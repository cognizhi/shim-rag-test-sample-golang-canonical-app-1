package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/calculate?op=add&left=7&right=5", nil)
	response := httptest.NewRecorder()

	NewHandler().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", response.Code, http.StatusOK, response.Body.String())
	}

	var body calculateResponse
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if body.Result != 12 {
		t.Fatalf("result = %v, want 12", body.Result)
	}
}

func TestCalculateRejectsUnknownOperation(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/calculate?op=pow&left=2&right=8", nil)
	response := httptest.NewRecorder()

	NewHandler().ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusBadRequest)
	}
}

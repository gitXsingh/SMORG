package home

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerRendersHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	NewHandler().ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.Code)
	}
	if !strings.Contains(res.Body.String(), "Keep everyday life moving") {
		t.Fatal("expected home page heading")
	}
}

func TestHandlerRejectsUnknownPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	res := httptest.NewRecorder()
	NewHandler().ServeHTTP(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", res.Code)
	}
}

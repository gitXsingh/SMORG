package capture

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHandlerAcceptsCapture(t *testing.T) {
	form := url.Values{"text": {"Finish the DBMS assignment"}}
	req := httptest.NewRequest(http.MethodPost, "/capture", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res := httptest.NewRecorder()

	NewHandler().ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.Code)
	}
	if !strings.Contains(res.Body.String(), "Finish the DBMS assignment") {
		t.Fatal("expected captured text in response")
	}
}

func TestHandlerRejectsEmptyCapture(t *testing.T) {
	form := url.Values{"text": {"   "}}
	req := httptest.NewRequest(http.MethodPost, "/capture", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res := httptest.NewRecorder()

	NewHandler().ServeHTTP(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", res.Code)
	}
}

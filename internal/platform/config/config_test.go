package config

import "testing"

func TestLoadUsesDefaultAddress(t *testing.T) {
	t.Setenv("APP_ADDRESS", "")
	if got := Load().Address; got != ":8080" {
		t.Fatalf("expected default address :8080, got %q", got)
	}
}

func TestLoadUsesConfiguredAddress(t *testing.T) {
	t.Setenv("APP_ADDRESS", ":9090")
	if got := Load().Address; got != ":9090" {
		t.Fatalf("expected configured address :9090, got %q", got)
	}
}

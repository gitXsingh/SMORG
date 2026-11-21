package main

import (
	"log"
	"net/http"

	"organizer/internal/features/capture"
	"organizer/internal/features/home"
	"organizer/internal/platform/config"
	"organizer/internal/platform/httpserver"
)

func main() {
	cfg := config.Load()
	handler := httpserver.New(home.NewHandler(), capture.NewHandler())
	server := &http.Server{
		Addr:              cfg.Address,
		Handler:           handler,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	log.Printf("organizer server listening on %s", cfg.Address)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

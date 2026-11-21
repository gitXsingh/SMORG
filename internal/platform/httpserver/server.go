package httpserver

import "net/http"

func New(homeHandler http.Handler, captureHandler http.Handler) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /", homeHandler)
	mux.Handle("POST /capture", captureHandler)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	return mux
}

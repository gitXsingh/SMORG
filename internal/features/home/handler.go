package home

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed home.html
var templateFiles embed.FS

type Handler struct {
	template *template.Template
}

func NewHandler() *Handler {
	return &Handler{template: template.Must(template.ParseFS(templateFiles, "home.html"))}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.template.Execute(w, map[string]string{"Title": "Organizer"}); err != nil {
		http.Error(w, "template rendering failed", http.StatusInternalServerError)
	}
}

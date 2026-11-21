package capture

import (
	"html/template"
	"net/http"
	"strings"
)

type Handler struct {
	template *template.Template
}

func NewHandler() *Handler {
	return &Handler{template: template.Must(template.New("capture").Parse(`<!doctype html>
<html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Captured</title></head>
<body style="font-family:system-ui,sans-serif;max-width:680px;margin:64px auto;padding:0 20px;color:#172033">
<h1>Capture received</h1><p>{{.Text}}</p><p><a href="/">Back to organizer</a></p>
</body></html>`))}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	text := strings.TrimSpace(r.FormValue("text"))
	if text == "" {
		http.Error(w, "capture text is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.template.Execute(w, map[string]string{"Text": text}); err != nil {
		http.Error(w, "capture rendering failed", http.StatusInternalServerError)
	}
}

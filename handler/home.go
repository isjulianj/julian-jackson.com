package handler

import (
	"html/template"
	"log/slog"
	"net/http"
)

func HomeHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./ui/components/document.tmpl",
			"./ui/components/header.tmpl",
			"./ui/components/footer.tmpl",
			"./ui/html/pages/home.tmpl",
		}

		t, err := template.ParseFiles(files...)
		if err != nil {
			logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = t.ExecuteTemplate(w, "document", nil)

		if err != nil {
			logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

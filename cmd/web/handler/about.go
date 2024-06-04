package handler

import (
	"github.com/isjulianj/gojulianJackson/cmd/web/config"
	"html/template"
	"net/http"
)

func AboutHandler(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./ui/components/document.tmpl",
			"./ui/components/header.tmpl",
			"./ui/components/footer.tmpl",
			"./ui/html/pages/about.tmpl",
		}

		t, err := template.ParseFiles(files...)
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		err = t.ExecuteTemplate(w, "document", nil)
		if err != nil {
			app.ServerError(w, r, err)
		}
	}

}

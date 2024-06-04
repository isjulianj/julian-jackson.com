package main

import (
	"embed"
	"github.com/isjulianj/gojulianJackson/config"
	"net/http"
)

func staticAssets(app *config.Application) http.FileSystem {
	return http.Dir(app.Config.AssetsDirectory)
}

//go:embed ui/*
var resources embed.FS

//var t = template.Must(template.ParseFS(resources, "ui/home/home.tmpl", "ui/components/*"))
//
//func (app *application) home(w http.ResponseWriter, r *http.Request) {
//
//	files := []string{
//		"./ui/components/document.tmpl",
//		"./ui/components/header.tmpl",
//		"./ui/components/footer.tmpl",
//		"./ui/home/home.tmpl",
//	}
//
//	t, err := template.ParseFiles(files...)
//	if err != nil {
//		app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		return
//	}
//
//	err = t.ExecuteTemplate(w, "document", nil)
//
//	if err != nil {
//		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//	}
//}
//
//func (app *application) about(w http.ResponseWriter, r *http.Request) {
//	files := []string{
//		"./ui/components/document.tmpl",
//		"./ui/components/header.tmpl",
//		"./ui/components/footer.tmpl",
//		"./ui/about/home.tmpl",
//	}
//
//	t, err := template.ParseFiles(files...)
//	if err != nil {
//		app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		return
//	}
//
//	err = t.ExecuteTemplate(w, "document", nil)
//	if err != nil {
//		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//	}
//}

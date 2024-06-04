package main

import (
	"fmt"
	"github.com/isjulianj/gojulianJackson/cmd/web/config"
	handler2 "github.com/isjulianj/gojulianJackson/cmd/web/handler"
	"net/http"
)

func Handler(app *config.Application) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(fmt.Sprintf("./%s", app.Config.AssetsDirectory)))
	mux.Handle("GET /assets/", http.StripPrefix(fmt.Sprintf("/%s", app.Config.AssetsDirectory), fileServer))
	mux.HandleFunc("GET /", handler2.HomeHandler(app))
	mux.HandleFunc("GET /about", handler2.AboutHandler(app))

	return mux
}

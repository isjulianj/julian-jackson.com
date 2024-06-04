package routes

import (
	"fmt"
	"github.com/isjulianj/gojulianJackson/config"
	"github.com/isjulianj/gojulianJackson/handler"
	"net/http"
)

func Handler(app *config.Application) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(fmt.Sprintf("./%s", app.Config.AssetsDirectory)))
	mux.Handle("GET /assets/", http.StripPrefix(fmt.Sprintf("/%s", app.Config.AssetsDirectory), fileServer))
	mux.HandleFunc("GET /", handler.HomeHandler(app.Logger))
	mux.HandleFunc("GET /about", handler.AboutHandler(app.Logger))

	return mux
}

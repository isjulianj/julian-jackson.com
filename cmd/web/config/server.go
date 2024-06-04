package config

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *Application) ServeHTTP(handler http.Handler) error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.HttpPort),
		Handler:      handler,
		ErrorLog:     slog.NewLogLogger(app.Logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.Logger.Info("started server", slog.Group("server", "addr", server.Addr))
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil

}

func (app *Application) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.Logger.Error(err.Error(), http.StatusText(http.StatusInternalServerError), slog.String("method", method), slog.String("uri", uri), slog.String("trace", trace))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

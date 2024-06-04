package config

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
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

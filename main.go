package main

import (
	"flag"
	"github.com/isjulianj/gojulianJackson/config"
	"github.com/isjulianj/gojulianJackson/routes"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := run(logger)
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func run(logger *slog.Logger) error {
	var cfg config.Config

	flag.IntVar(&cfg.HttpPort, "port", 8080, "http server port")
	flag.StringVar(&cfg.Env, "env", "development", "development|staging|production")
	flag.StringVar(&cfg.AssetsDirectory, "assetDir", "assets", "assets directory")
	flag.Parse()

	app := config.NewApplication(cfg, logger)

	muxHandler := routes.Handler(app)

	return app.ServeHTTP(muxHandler)

}

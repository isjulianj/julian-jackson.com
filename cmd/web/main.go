package main

import (
	"flag"
	"github.com/isjulianj/gojulianJackson/cmd/web/config"
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

	muxHandler := Handler(app)

	return app.ServeHTTP(muxHandler)

}

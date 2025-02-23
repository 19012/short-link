package main

import (
	"19012/short-link/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// export CONFIG_PATH=~/go/short-link/config/local.yaml
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("short-link: launch", slog.String("env", cfg.Env))
	log.Debug("debug messages are anabled")

	// TODO: logger(slog)
	// TODO: storage(sqlite)
	// TODO: router(chi, chi render)
	// TODO: run service
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

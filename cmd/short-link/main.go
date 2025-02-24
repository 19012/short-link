package main

import (
	"19012/short-link/internal/config"
	"19012/short-link/internal/lib/logger/sl"
	"19012/short-link/internal/storage/sqlite"
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

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		slog.Error("fail to init storage", sl.ErrAttr(err))
		os.Exit(1)
	}
	_ = storage

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

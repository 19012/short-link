package main

import (
	"19012/short-link/internal/config"
)

func main() {
	// TODO: init config (cleanenv)
	cfg := config.MustLoad()
	_ = cfg
	// TODO:init logger (log/slog)

	// TODO: init db (sqlite)

	// TODO: init router (chi, chi render)

	// TODO: launch server
}

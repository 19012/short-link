package main

import (
	"19012/short-link/internal/config"
	"fmt"
)

func main() {
	// export CONFIG_PATH=~/go/short-link/config/local.yaml
	cfg := config.MustLoad()
	fmt.Println(cfg) // TODO: delete this line
	// TODO: logger(slog)
	// TODO: storage(sqlite)
	// TODO: router(chi, chi render)
	// TODO: run service
}

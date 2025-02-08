package main

import (
	"19012/short-link/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// TODO: init logger
	// TODO: init storage
	// TODO: init router
	// TODO: run
}

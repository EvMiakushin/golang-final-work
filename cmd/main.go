package main

import (
	"golang-final-work/internal/config"
	"golang-final-work/internal/handlers"
)

func main() {
	cfg := config.GetConfig()
	srv := handlers.NewServer()
	srv.Start(cfg.Port)
}

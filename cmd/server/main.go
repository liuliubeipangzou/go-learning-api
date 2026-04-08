package main

import (
	"log"

	"go-learning-api/internal/app"
	"go-learning-api/internal/config"
)

func main() {
	cfg := config.Load()

	application := app.New(cfg)

	log.Printf("%s starting on :%s", cfg.AppName, cfg.Port)

	if err := application.Run(); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}

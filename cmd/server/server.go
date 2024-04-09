package main

import (
	"github.com/tasticolly/gRPC-chat/internal/config"
	"log"
	"log/slog"
	"os"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}

func main() {
	cfg, err := config.Load()
	log.Fatal(err)

	logger := setupLogger(cfg.Env)

}

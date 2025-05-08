package main

import (
	"context"
	"news-mono/cmd/internal/app"
	"news-mono/cmd/internal/config"
	"news-mono/cmd/pkg/logging"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logging.GetLogger(ctx)
	logger.Info("config initializing")

	cfg := config.GetConfig()

	ctx = logging.ContextWithLogger(ctx, logger)

	a, err := app.NewApp(ctx, cfg)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Running Application")

	a.Run(ctx)
}

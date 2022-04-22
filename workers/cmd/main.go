package main

import (
	"github/Services/workers/api"
	"github/Services/workers/config"
	"github/Services/workers/pkg/logger"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api-gateway")

	

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}

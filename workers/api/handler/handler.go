package handlers

import (
	"github/Services/workers/config"
	"github/Services/workers/pkg/logger"
	"github/Services/workers/storage/repo"
)

type handlerV1 struct {
	log             logger.Logger
	cfg             config.Config
	inMemoryStorage repo.UserStorageI
}

type HandlerV1Config struct {
	Logger          logger.Logger
	Cfg             config.Config
	InMemoryStorage repo.UserStorageI
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:             c.Logger,
		cfg:             c.Cfg,
		inMemoryStorage: c.InMemoryStorage,
	}
}

package handlers

import (
	"github/Services/workers/config"
	"github/Services/workers/storage/repo"
)

type handlerV1 struct {
	cfg             config.Config
	inMemoryStorage repo.UserStorageI
}

type HandlerV1Config struct {
	Cfg             config.Config
	InMemoryStorage repo.UserStorageI
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		cfg:             c.Cfg,
		inMemoryStorage: c.InMemoryStorage,
	}
}

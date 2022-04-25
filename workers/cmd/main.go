package main

import (
	"fmt"
	"github/Services/workers/api"
	"github/Services/workers/config"
	"github/Services/workers/storage"
	_ "github/Services/workers/storage/postgres"

	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.Load()

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		panic(err)
	}

	storagePool := storage.NewStoragePg(connDb).User()

	server := api.New(api.Option{
		Conf:            cfg,
		InMemoryStorage: storagePool,
	})

	err = server.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}

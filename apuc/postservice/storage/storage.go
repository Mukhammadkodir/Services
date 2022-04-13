package storage

import (
	"github/Services/apuc/postservice/storage/postgres"
	"github/Services/apuc/postservice/storage/repo"

	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
	Post() repo.PostStorageI
}

type storagePg struct {
	db       *sqlx.DB
	PostRepo repo.PostStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		PostRepo: postgres.NewPostRepo(db),
	}
}

func (s storagePg) Post() repo.PostStorageI {
	return s.PostRepo
}

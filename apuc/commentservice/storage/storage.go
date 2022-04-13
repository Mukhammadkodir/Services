package storage

import (
	"github/Services/apuc/commentservice/storage/postgres"
	"github/Services/apuc/commentservice/storage/repo"

	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
	Comment() repo.CommentStorageI
}

type storagePg struct {
	db       *sqlx.DB
	commentRepo repo.CommentStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		commentRepo: postgres.NewCommentRepo(db),
	}
}

func (s storagePg) Comment() repo.CommentStorageI {
	return s.commentRepo
}

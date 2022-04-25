package repo

import (
	"github/Services/workers/storage/models"
)

//UserStorageI ...
type UserStorageI interface {
	Create(m models.User) (*models.User, error)
	Get(id models.ById) (*models.Get, error)
	Update(new models.UpReq) (*models.User, error)
	Delete(pass models.PasswordReq) (*models.EmptyResp, error)
	Login(pass models.PasswordReq) (*models.User, error)
	ListUser(L models.ListReq) (*models.ListResp, error)
	CheckField(pass models.PasswordReq) (*models.Status, error)
	OpenDay(id models.ById) (*models.Hour, error)
}

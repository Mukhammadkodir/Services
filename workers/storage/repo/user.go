package repo

import (
	pb "github/Services/workers/genproto/user_service"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	Get(*pb.PasswordReq) (*pb.GetUser, error)
	Update(*pb.UpReq) (*pb.User, error)
	Delete(*pb.PasswordReq) (*pb.EmptyResp, error)
	Login(*pb.PasswordReq) (*pb.User, error)
	ListUser(*pb.ListReq)(*pb.ListResp, error)
	CheckField(*pb.PasswordReq)(*pb.Status,error)
	OpenDay(*pb.ById) (*pb.Hours,error)
	//ListTime(*pb.ById) (*pb.Status, error)
}

package repo

import (
	pb "github/Services/apuc/userservice/genproto/user_service"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	Get(*pb.ById) (*pb.User, error)
	Delete(*pb.ById) (*pb.Empty, error)
	Update(*pb.User) (*pb.User, error)
	ListUser(*pb.ListReq)(*pb.ListResp, error)
}

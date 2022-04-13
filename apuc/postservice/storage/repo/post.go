package repo

import (
	pb "github/Services/apuc/postservice/genproto/post_service"
)

//PostStorageI ...
type PostStorageI interface {
	Create(*pb.Post) (*pb.Post, error)
	Get(*pb.ById) (*pb.Post, error)
	Delete(*pb.ById) (*pb.Empty, error)
	Update(*pb.Post) (*pb.Post, error)
	GetById(*pb.ByUId) (*pb.ListResp, error)
	List(*pb.ListReq) (*pb.ListResp, error)
}

package repo

import (
	pb "github/Services/apuc/commentservice/genproto/comment_service"
)

//CommentStorageI ...
type CommentStorageI interface {
	Create(*pb.Comment) (*pb.Comment, error)
	Get(*pb.ById) (*pb.Comment, error)
	Delete(*pb.ById) (*pb.Empty, error)
	Update(*pb.Comment) (*pb.Comment, error)
	ListComment(*pb.ListReq)(*pb.ListResp, error)
	DeleteByUser(*pb.ById) (*pb.Empty, error)
}

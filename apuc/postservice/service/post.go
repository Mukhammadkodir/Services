package service

import (
	"context"

	pb "github/Services/apuc/postservice/genproto/post_service"
	l "github/Services/apuc/postservice/pkg/logger"

	"github/Services/apuc/postservice/storage"

	"github.com/jmoiron/sqlx"
)

//PostService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	user, err := s.storage.Post().Create(req)
	if err != nil {
		s.logger.Error("Failed create post", l.Error(err))
		return nil, nil
	}

	return user, nil
}

func (s *PostService) Delete(ctx context.Context, req *pb.ById) (*pb.Empty, error) {

	_, err := s.storage.Post().Delete(req)
	if err != nil {
		s.logger.Error("Failed Delete post", l.Error(err))
		return nil, nil
	}

	return &pb.Empty{}, nil
}

func (s *PostService) Update(ctx context.Context, req *pb.Post) (*pb.Post, error) {

	post, err := s.storage.Post().Update(req)
	if err != nil {
		s.logger.Error("Failed update post", l.Error(err))
		return nil, nil
	}

	return post, nil
}

func (s *PostService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {

	resp, err := s.storage.Post().List(req)
	if err != nil {
		s.logger.Error("Error getting list Posts", l.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *PostService) GetById(ctx context.Context, req *pb.ByUId) (*pb.ListResp, error) {

	posts, err := s.storage.Post().GetById(req)
	if err != nil {
		s.logger.Error("Failed to get password with id", l.Error(err))
		return nil, nil
	}

	return posts, nil
}

func (s *PostService) Get(ctx context.Context, req *pb.ById) (*pb.Post, error) {

	post, err := s.storage.Post().Get(req)
	if err != nil {
		s.logger.Error("Failed to get post with id", l.Error(err))
		return nil, nil
	}

	return post, nil
}

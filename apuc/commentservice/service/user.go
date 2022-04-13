package service

import (
	"context"

	pb "github/Services/apuc/commentservice/genproto/comment_service"
	l  "github/Services/apuc/commentservice/pkg/logger"
	   "github/Services/apuc/commentservice/storage"

	"github.com/jmoiron/sqlx"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
	user, err := s.storage.Comment().Create(req)
	if err != nil {
		s.logger.Error("Failed create user", l.Error(err))
		return nil, nil
	}

	return user, nil
}

func (s *UserService) Get(ctx context.Context, req *pb.ById) (*pb.Comment, error) {
	user, err := s.storage.Comment().Get(req)
	if err != nil {
		s.logger.Error("Failed Get user", l.Error(err))
		return nil, nil
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.ById) (*pb.Empty, error) {
	_, err := s.storage.Comment().Delete(req)
	if err != nil {
		s.logger.Error("Failed Delete user", l.Error(err))
		return nil, nil
	}

	return &pb.Empty{}, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
	user, err := s.storage.Comment().Update(req)
	if err != nil {
		s.logger.Error("Failed update user", l.Error(err))
		return nil, nil
	}

	return user, nil
}

func (s *UserService) ListUser(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
	resp, err := s.storage.Comment().ListComment(req)
	if err != nil {
		s.logger.Error("Error list Users", l.Error(err))
		return nil, err
	}
	return resp, nil
}

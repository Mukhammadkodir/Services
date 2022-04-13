package service

import (
	"context"

	pb "github/Services/apuc/userservice/genproto/user_service"
	l "github/Services/apuc/userservice/pkg/logger"
	grpcClient "github/Services/apuc/userservice/service/grpcclient"
	p "github/Services/apuc/userservice/genproto/post_service"

	"github/Services/apuc/userservice/storage"

	"github.com/jmoiron/sqlx"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.GrpcClientI
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("Failed create user", l.Error(err))
		return nil, nil
	}

	return user, nil
}

func (s *UserService) Get(ctx context.Context, req *pb.ById) (*pb.User, error) {
	user, err := s.storage.User().Get(req)
	if err != nil {
		s.logger.Error("Failed Get user", l.Error(err))
		return nil, nil
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.ById) (*pb.Empty, error) {
	_, err := s.storage.User().Delete(req)
	if err != nil {
		s.logger.Error("Failed Delete user", l.Error(err))
		return nil, nil
	}

	comment, err := s.client.PostService().DeleteByUser(context.Background(), &p.ById{Userid: req.Userid})

	if err != nil {
		s.logger.Error("Error delete post from comment_service", l.Error(err))
		return nil, err
	}

	return (*pb.Empty)(comment), nil
}

func (s *UserService) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().Update(req)
	if err != nil {
		s.logger.Error("Failed update user", l.Error(err))
		return nil, nil
	}

	return user, nil
}

func (s *UserService) ListUser(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
	resp, err := s.storage.User().ListUser(req)
	if err != nil {
		s.logger.Error("Error list Users", l.Error(err))
		return nil, err
	}
	return resp, nil
}

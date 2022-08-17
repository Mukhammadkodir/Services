package service

import (
	"context"
	"fmt"

	p "github/Services/apuc/userservice/genproto/post_service"
	pb "github/Services/apuc/userservice/genproto/user_service"
	l "github/Services/apuc/userservice/pkg/logger"
	"github/Services/apuc/userservice/pkg/messagebroker"
	grpcClient "github/Services/apuc/userservice/service/grpcclient"

	"github/Services/apuc/userservice/storage"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
	storage   storage.IStorage
	logger    l.Logger
	client    grpcClient.GrpcClientI
	publisher map[string]messagebroker.Publisher
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI, publisher map[string]messagebroker.Publisher) *UserService {
	return &UserService{
		storage:   storage.NewStoragePg(db),
		logger:    log,
		client:    client,
		publisher: publisher,
	}
}

func (s *UserService) publisherUserMessage(user []byte) error {

	err := s.publisher["user"].Publish([]byte("user"), user, string(user))
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("Failed create user", l.Error(err))
		return nil, nil
	}

	p, _ := user.Marshal()
	var usera pb.User
	err = usera.Unmarshal(p)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(user)

	err = s.publisherUserMessage(p)
	if err != nil {
		s.logger.Error("failed while publishing user info", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while publishing user info")

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

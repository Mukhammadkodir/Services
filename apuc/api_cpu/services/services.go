package services

import (
	"fmt"
	"github/Services/apuc/api_cpu/config"
	pb "github/Services/apuc/api_cpu/genproto/post_service"
	pbu "github/Services/apuc/api_cpu/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	PostService() pb.UserServiceClient
	UserService() pbu.UserServiceClient
}

type serviceManager struct {
	postService pb.UserServiceClient
	userService pbu.UserServiceClient
}

func (s *serviceManager) PostService() pb.UserServiceClient {
	return s.postService
}
func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}


func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	

	serviceManager := &serviceManager{
		postService: pb.NewUserServiceClient(connPost),
		userService: pbu.NewUserServiceClient(connUser),
	}

	return serviceManager, nil
}

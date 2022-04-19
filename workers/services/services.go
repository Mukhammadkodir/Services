package services

import (
	"fmt"
	"github/Services/workers/config"
	pbu "github/Services/workers/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
}

type serviceManager struct {
	userService pbu.UserServiceClient
}


func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	

	serviceManager := &serviceManager{
		userService: pbu.NewUserServiceClient(connUser),
	}

	return serviceManager, nil
}

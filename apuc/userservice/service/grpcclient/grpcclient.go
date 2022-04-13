package grpc_client

import (
	"fmt"
	"github/Services/apuc/userservice/config"

	pb "github/Services/apuc/userservice/genproto/post_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	PostService() pb.UserServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	postService pb.UserServiceClient
}

func New(cfg config.Config) (*GrpcClient, error) {

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port: %d",
			cfg.PostServiceHost, cfg.PostServicePort)
	}

	grpcClient := &GrpcClient{
		cfg:         cfg,
		postService: pb.NewUserServiceClient(connPost),
	}

	return grpcClient, nil
}

// TaskService ...
func (c *GrpcClient) PostService() pb.UserServiceClient {
	return c.postService
}

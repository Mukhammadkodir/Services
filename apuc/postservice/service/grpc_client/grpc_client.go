package grpc_client

import (
	"fmt"
	"github/Services/apuc/postservice/config"

	pb "github/Services/apuc/postservice/genproto/comment_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	CommentService() pb.CommentServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	commentService pb.CommentServiceClient
}

func New(cfg config.Config) (*GrpcClient, error) {

	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port: %d",
			cfg.CommentServiceHost, cfg.CommentServicePort)
	}

	grpcClient := &GrpcClient{
		cfg:         cfg,
		commentService: pb.NewCommentServiceClient(connComment),
	}

	return grpcClient, nil
}

// TaskService ...
func (c *GrpcClient) CommentService() pb.CommentServiceClient {
	return c.commentService
}

package grpc_client

// import (
// 	"fmt"
// 	"github/Services/post_task/post_service/config"
// 	pb "github/Services/post_task/post_service/genproto/data_service"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// type GrpcClientI interface {
// 	DataService() pb.DataServiceClient
// }

// type GrpcClient struct {
// 	cfg         config.Config
// 	dataService pb.DataServiceClient
// }

// func New(cfg config.Config) (*GrpcClient, error) {
	
// 	connData, err := grpc.Dial(
// 		fmt.Sprintf("%s:%d", cfg.DataServiceHost, cfg.DataServicePort),
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	)

// 	if err != nil {
// 		return nil, fmt.Errorf("catalog service dial host: %s port: %d",
// 			cfg.DataServiceHost, cfg.DataServicePort)
// 	}

// 	grpcClient := &GrpcClient{
// 		cfg:         cfg,
// 		dataService: pb.NewDataServiceClient(connData),
// 	}

// 	return grpcClient, nil
// }

// // DataService ...
// func (c *GrpcClient) DataService() pb.DataServiceClient {
// 	return c.dataService
// }


package main

import (
	"github/Services/apuc/userservice/config"
	"github/Services/apuc/userservice/events"
	pb "github/Services/apuc/userservice/genproto/user_service"
	"github/Services/apuc/userservice/pkg/db"
	"github/Services/apuc/userservice/pkg/logger"
	"github/Services/apuc/userservice/pkg/messagebroker"
	"github/Services/apuc/userservice/service"
	grpc_client "github/Services/apuc/userservice/service/grpcclient"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "User")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
		logger.String("password", cfg.PostgresPassword))
	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	grpcClient, err := grpc_client.New(cfg)
	if err != nil {
		log.Error("error establishing grpc connection", logger.Error(err))
		return
	}

	//Kafka
	publisherMap := make(map[string]messagebroker.Publisher)

	userTopicPublisher := events.NewKafkaProducerBroker(cfg, log, "user.user")
	defer func() {
		err := userTopicPublisher.Stop()
		if err != nil {
			log.Fatal("failed to stop kafka user", logger.Error(err))
		}
	}()

	publisherMap["user"] = userTopicPublisher
	//Kafka End

	userService := service.NewUserService(connDB, log, grpcClient, publisherMap)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}

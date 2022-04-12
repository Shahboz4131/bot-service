package main

import (
	"net"

	pb "github.com/Shahboz4131/bot-service/genproto"
	"github.com/Shahboz4131/bot-service/pkg/logger"
	"github.com/Shahboz4131/bot-service/service"

	"github.com/Shahboz4131/bot-service/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "bot-service")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	botService := service.NewBotService(log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterBotServiceServer(s, botService)
	reflection.Register(s)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}

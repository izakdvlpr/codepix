package grpc

import (
	"fmt"
	"github.com/izakdvlpr/codepix/application/grpc/pb"
	"github.com/izakdvlpr/codepix/application/usecase"
	"github.com/izakdvlpr/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	pixKeyRepository := repository.PixKeyRepositoryDatabase{Database: database}
	pixKeyUseCase := usecase.PixKeyUseCase{PixKeyRepository: pixKeyRepository}
	pixKeyGrpcService := NewPixKeyGrpcService(pixKeyUseCase)

	pb.RegisterPixServiceServer(grpcServer, pixKeyGrpcService)

	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}

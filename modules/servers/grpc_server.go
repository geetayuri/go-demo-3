package servers

import (
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	pb "github.com/nguitarpb/7-solutions/protobuf/proto"
	_searchlistGrpc "github.com/nguitarpb/7-solutions/modules/searchlist/grpc_controllers"
	_searchlistRepository "github.com/nguitarpb/7-solutions/modules/searchlist/repositories"
	_searchlistUsecase "github.com/nguitarpb/7-solutions/modules/searchlist/usecases"
)

const (
	port = ":6000"
)

type GrpcServer struct {
	pb.UnimplementedDataServiceServer
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (gs *GrpcServer) StartGRPCServer(wg *sync.WaitGroup) {
	defer wg.Done()

	// Initialize repositories, usecases, and controllers
	searchlistRepository := _searchlistRepository.NewSearchlistRepository()
	searchlistUsecase := _searchlistUsecase.NewSearchlistUsecase(searchlistRepository)
	controller := _searchlistGrpc.NewSearchlistController(searchlistUsecase)

	// Start gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// Register the controller with the gRPC server
	pb.RegisterDataServiceServer(grpcServer, controller)

	log.Printf("gRPC Server is listening on port %v", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
package grpc_controllers

import (
	"context"
	// "fmt"

	"github.com/nguitarpb/7-solutions/modules/searchlist/entities"
	pb "github.com/nguitarpb/7-solutions/protobuf/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type searchlistGrpcCon struct {
	SearchlistUse entities.SearchlistUsecase
	pb.UnimplementedDataServiceServer
}

func NewSearchlistController(searchlistUse entities.SearchlistUsecase) *searchlistGrpcCon {
	return &searchlistGrpcCon{
		SearchlistUse: searchlistUse,
	}
}

func (h *searchlistGrpcCon) GetData(ctx context.Context, in *pb.Empty) (*pb.DataResponse, error) {
	res, err := h.SearchlistUse.Search()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	// fmt.Println(res.Beef)

	// return nil, status.Errorf(codes.NotFound, "Internal Server Error")
	return &pb.DataResponse{Beef: res.Beef}, nil
}

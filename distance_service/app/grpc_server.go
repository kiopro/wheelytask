package app

import (
	"context"
	"net"
	"os"

	pb "../proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	gRPCPort = ":8222"
)

type gRPCServer struct{}

func (s *gRPCServer) GetTripInformation(ctx context.Context, trip *pb.Trip) (*pb.TripInformationResponse, error) {
	element, err := DistanceMatrixRequest(trip)
	if err != nil {
		return nil, err
	}

	distance := element.Distance["value"]
	time := element.Duration["value"]

	return &pb.TripInformationResponse{Distance: distance, Time: time}, nil
}

// StartGRPCServer : ...
func StartGRPCServer() {
	gRPCPort := os.Getenv("GRPC_PORT")

	lis, err := net.Listen("tcp", gRPCPort)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTripInformationServer(s, &gRPCServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}

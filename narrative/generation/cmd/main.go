package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "generation/pkg/proto/generation"

	"github.com/gage-technologies/mistral-go"
)

type GenerationServiceServer struct {
	generation.UnimplementedGenerationServiceServer
	Client *mistral.NewMistralClientDefault
	ModelId string
}

func (s *GenerationServiceServer) GenerateResponseStream(req *pb.GenerateResponseRequest, stream pb.GenerationService_GenerateResponseStreamServer) error {
	log.Printf("Received request: %v", req)
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.StreamResponse{Token: "dummy token"}); err != nil {
			return err
		}
	}
	return nil
}

func (s *GenerationServiceServer) GenerateResponseBatch(context.Context, *pb.GenerateResponseRequest) (*pb.BatchResponse, error) {
	return &pb.BatchResponse{Response: "dummy response"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	generation.RegisterGenerationServiceServer(grpcServer, &GenerationServiceServer{})
	log.Printf("Server is running on port 50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

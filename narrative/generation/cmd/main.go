package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	pb "generation/pkg/proto/generation"
	"generation/pkg/dummy"
	"github.com/gage-technologies/mistral-go"
	"generation/pkg/mistral"
)

type GenerationServiceServer struct {
	pb.UnimplementedGenerationServiceServer
	Generator generation.Generator
}

func (s *GenerationServiceServer) GenerateResponseStream(req *pb.GenerateResponseRequest, stream pb.GenerationService_GenerateResponseStreamServer) error {
	return s.Generator.GenerateResponseStream(req, stream)
}

func (s *GenerationServiceServer) GenerateResponseBatch(ctx context.Context, req *pb.GenerateResponseRequest) (*pb.BatchResponse, error) {
	return s.Generator.GenerateResponseBatch(ctx, req)
}

func main() {
	generatorType := flag.String("generator", "dummy", "The type of generator to use (dummy, mistral)")
	flag.Parse()

	var generator generation.Generator

	switch *generatorType {
	case "dummy":
		generator = dummy.NewDummyGenerator()
	case "mistral":
		mistralClient := mistral.NewMistralClientDefault()
		modelId := os.Getenv("MISTRAL_MODEL_ID")
		generator = mistral.NewMistralGenerator(mistralClient, modelId)
	default:
		log.Fatalf("Unsupported generator type: %s", *generatorType)
	}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGenerationServiceServer(grpcServer, &GenerationServiceServer{Generator: generator})
	log.Printf("Server is running on port 50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

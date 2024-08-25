package generation

import (
	"context"
	pb "generation/pkg/proto/generation"
)

type Generator interface {
	GenerateResponseStream(req *pb.GenerateResponseRequest, stream pb.GenerationService_GenerateResponseStreamServer) error
	GenerateResponseBatch(ctx context.Context, req *pb.GenerateResponseRequest) (*pb.BatchResponse, error)
}

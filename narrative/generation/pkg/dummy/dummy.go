import (
	"context"
	pb "generation/pkg/proto/generation"
	"log"
)

type DummyGenerator struct{}

func (d *DummyGenerator) GenerateResponseStream(req *pb.GenerateResponseRequest, stream pb.GenerationService_GenerateResponseStreamServer) error {
	log.Printf("Received request: %v", req)
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.StreamResponse{Token: "dummy token"}); err != nil {
			return err
		}
	}
	return nil
}

func (d *DummyGenerator) GenerateResponseBatch(ctx context.Context, req *pb.GenerateResponseRequest) (*pb.BatchResponse, error) {
	log.Printf("Recieved request: %v", req)
	return &pb.BatchResponse{Response: "dummy response"}, nil
}

func NewDummyGenerator() *DummyGenerator {
	return &DummyGenerator{}
}

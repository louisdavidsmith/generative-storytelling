package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "query/pkg/proto/query"
	"github.com/jackc/pgx/v5/pgxpool"
)

type server struct {
	pb.UnimplementedQueryServiceServer
	Pool *pgxpool.Pool
}

func (s *server) CreateAdventure(ctx context.Context, req *pb.CreateAdventureRequest) (*pb.CreateAdventureResponse, error) {
	sql := `INSERT INTO adventure (adventure_id, lore_id, premise) VALUES ($1, $2, $3)`
	_, err = pool.Exec(context.Background(), stmt, req.AdventureId, req.LoreId, req.Premise, req.AdventureName)
	if err != nil {
		log.Fatal(err)
	}

	res := &pb.CreateAdventureResponse{
	}
	return res, nil
}

func (s *server) CreateAdventureLore(ctx context.Context, req *pb.CreateAdventureLoreRequest) (*pb.CreateAdventureLoreResponse, error) {
	// Implement the logic for creating adventure lore here.
	res := &pb.CreateAdventureLoreResponse{
		LoreName: req.LoreName,
		LoreId:   "67890", // Placeholder ID
	}
	return res, nil
}

func (s *server) CreateCharacter(ctx context.Context, req *pb.CreateCharacterRequest) (*pb.CreateCharacterResponse, error) {
	// Implement the logic for creating a character here.
	res := &pb.CreateCharacterResponse{
		CharacterName: req.CharacterName,
		CharacterId:   "abcde", // Placeholder ID
	}
	return res, nil
}

func (s *server) GetAdventureLore(ctx context.Context, req *pb.GetAdventureLoreRequest) (*pb.GetAdventureLoreResponse, error) {
	// Implement the logic for retrieving adventure lore here.
	res := &pb.GetAdventureLoreResponse{
		LoreId:  req.LoreId,
		Context: "Some context about the lore", // Placeholder context
	}
	return res, nil
}

func (s *server) GetConversationHistory(ctx context.Context, req *pb.GetConversationHistoryRequest) (*pb.GetConversationHistoryResponse, error) {
	// Implement the logic for retrieving conversation history here.
	history := []*pb.Chat{
		{Content: "Hello, this is a message.", Role: pb.Role_USER},
		{Content: "Hello, this is the assistant.", Role: pb.Role_ASSISTANT},
	}
	res := &pb.GetConversationHistoryResponse{
		AdventureId: req.AdventureId,
		History:     history,
	}
	return res, nil
}

func (s *server) GetPlayerCharacteristics(ctx context.Context, req *pb.GetPlayerCharacteristicsRequest) (*pb.GetPlayerCharacteristicsResponse, error) {
	// Implement the logic for retrieving player characteristics here.
	stats := &pb.CharacterStats{
		Sanity:    100,
		Physical:  100,
		Wit:       100,
		Composure: 100,
		Social:    100,
	}
	res := &pb.GetPlayerCharacteristicsResponse{
		CharacterName: req.CharacterId,
		Statistics:    stats,
	}
	return res, nil
}

func (s *server) GetNarrativeState(ctx context.Context, req *pb.GetNarrativeStateRequest) (*pb.GetNarrativeStateResponse, error) {
	// Implement the logic for retrieving the narrative state here.
	res := &pb.GetNarrativeStateResponse{
		NarrativeState: "This is the current narrative state",
	}
	return res, nil
}

func (s *server) UpdateConversationHistory(ctx context.Context, req *pb.UpdateConversationHistoryRequest) (*pb.UpdateConversationHistoryResponse, error) {
	// Implement the logic for updating conversation history here.
	res := &pb.UpdateConversationHistoryResponse{
		AdventureId: req.AdventureId,
	}
	return res, nil
}

func (s *server) UpdatePlayerCharacteristics(ctx context.Context, req *pb.UpdatePlayerCharacteristicsRequest) (*pb.UpdatePlayerCharacteristicsResponse, error) {
	res := &pb.UpdatePlayerCharacteristicsResponse{
		AdventureId: req.AdventureId,
		CharacterId: req.CharacterId,
	}
	return res, nil
}

func (s *server) UpdateNarrativeState(ctx context.Context, req *pb.UpdateNarrativeStateRequest) (*pb.UpdateNarrativeStateResponse, error) {
	res := &pb.UpdateNarrativeStateResponse{
		AdventureId: req.AdventureId,
	}
	return res, nil
}

func main() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterQueryServiceServer(s, &server{dbpool})

	fmt.Println("Server is running on port 50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

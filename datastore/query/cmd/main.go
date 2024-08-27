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
	generation.UnimplementedQueryServiceServer
	Pool *pgxpool.Pool
}

func (s *server) CreateAdventure(ctx context.Context, req *generation.CreateAdventureRequest) (*generation.CreateAdventureResponse, error) {
	// Implement the logic for creating an adventure here.
	res := &generation.CreateAdventureResponse{
		AdventureName: req.AdventureName,
		AdventureId:   "12345", // Placeholder ID
	}
	return res, nil
}

func (s *server) CreateAdventureLore(ctx context.Context, req *generation.CreateAdventureLoreRequest) (*generation.CreateAdventureLoreResponse, error) {
	// Implement the logic for creating adventure lore here.
	res := &generation.CreateAdventureLoreResponse{
		LoreName: req.LoreName,
		LoreId:   "67890", // Placeholder ID
	}
	return res, nil
}

func (s *server) CreateCharacter(ctx context.Context, req *generation.CreateCharacterRequest) (*generation.CreateCharacterResponse, error) {
	// Implement the logic for creating a character here.
	res := &generation.CreateCharacterResponse{
		CharacterName: req.CharacterName,
		CharacterId:   "abcde", // Placeholder ID
	}
	return res, nil
}

func (s *server) GetAdventureLore(ctx context.Context, req *generation.GetAdventureLoreRequest) (*generation.GetAdventureLoreResponse, error) {
	// Implement the logic for retrieving adventure lore here.
	res := &generation.GetAdventureLoreResponse{
		LoreId:  req.LoreId,
		Context: "Some context about the lore", // Placeholder context
	}
	return res, nil
}

func (s *server) GetConversationHistory(ctx context.Context, req *generation.GetConversationHistoryRequest) (*generation.GetConversationHistoryResponse, error) {
	// Implement the logic for retrieving conversation history here.
	history := []*generation.Chat{
		{Content: "Hello, this is a message.", Role: generation.Role_USER},
		{Content: "Hello, this is the assistant.", Role: generation.Role_ASSISTANT},
	}
	res := &generation.GetConversationHistoryResponse{
		AdventureId: req.AdventureId,
		History:     history,
	}
	return res, nil
}

func (s *server) GetPlayerCharacteristics(ctx context.Context, req *generation.GetPlayerCharacteristicsRequest) (*generation.GetPlayerCharacteristicsResponse, error) {
	// Implement the logic for retrieving player characteristics here.
	stats := &generation.CharacterStats{
		Sanity:    100,
		Physical:  100,
		Wit:       100,
		Composure: 100,
		Social:    100,
	}
	res := &generation.GetPlayerCharacteristicsResponse{
		CharacterName: req.CharacterId,
		Statistics:    stats,
	}
	return res, nil
}

func (s *server) GetNarrativeState(ctx context.Context, req *generation.GetNarrativeStateRequest) (*generation.GetNarrativeStateResponse, error) {
	// Implement the logic for retrieving the narrative state here.
	res := &generation.GetNarrativeStateResponse{
		NarrativeState: "This is the current narrative state",
	}
	return res, nil
}

func (s *server) UpdateConversationHistory(ctx context.Context, req *generation.UpdateConversationHistoryRequest) (*generation.UpdateConversationHistoryResponse, error) {
	// Implement the logic for updating conversation history here.
	res := &generation.UpdateConversationHistoryResponse{
		AdventureId: req.AdventureId,
	}
	return res, nil
}

func (s *server) UpdatePlayerCharacteristics(ctx context.Context, req *generation.UpdatePlayerCharacteristicsRequest) (*generation.UpdatePlayerCharacteristicsResponse, error) {
	res := &generation.UpdatePlayerCharacteristicsResponse{
		AdventureId: req.AdventureId,
		CharacterId: req.CharacterId,
	}
	return res, nil
}

func (s *server) UpdateNarrativeState(ctx context.Context, req *generation.UpdateNarrativeStateRequest) (*generation.UpdateNarrativeStateResponse, error) {
	res := &generation.UpdateNarrativeStateResponse{
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
	generation.RegisterQueryServiceServer(s, &server{dbpool})

	fmt.Println("Server is running on port 50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

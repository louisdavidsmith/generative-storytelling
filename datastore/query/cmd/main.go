package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	pb "query/pkg/proto/query"
)

type server struct {
	pb.UnimplementedQueryServiceServer
	Pool *pgxpool.Pool
}

func (s *server) CreateAdventure(ctx context.Context, req *pb.CreateAdventureRequest) (*pb.CreateAdventureResponse, error) {
	sql := `INSERT INTO adventure (adventure_id, lore_id, premise, adventure_name) VALUES ($1, $2, $3, $4)`
	_, err = pool.Exec(context.Background(), sql, req.AdventureId, req.LoreId, req.Premise, req.AdventureName)
	if err != nil {
		log.Fatal(err)
	}

	res := &pb.CreateAdventureResponse{
		AdventureId: req.AdventureName,
	}
	return res, nil
}

func (s *server) CreateAdventureLore(ctx context.Context, req *pb.CreateAdventureLoreRequest) (*pb.CreateAdventureLoreResponse, error) {
	sql := `INSERT INTO lore (lore_id, lore_name, lore_description) VALUES ($1, $2, $3, $4)`
	_, err = pool.Exec(context.Background(), sql, req.LoreId, req.LoreName, req.LoreDescription)
	if err != nil {
		log.Fatal(err)
	}
	res := &pb.CreateAdventureLoreResponse{
		LoreName: req.LoreId,
	}
	return res, nil
}

func (s *server) CreateCharacter(ctx context.Context, req *pb.CreateCharacterRequest) (*pb.CreateCharacterResponse, error) {
	sql := `INSERT INTO game_character (game_id, character_id, name, description) VALUES ($1, $2, $3, $4)`
	_, err = pool.Exec(context.Background(), sql, req.GameId, req.CharacterId, req.Name, req.Description)
	if err != nil {
		log.Fatal(err)
	}
	statFields := []struct {
		name  string
		value int32
	}{
		{"sanity", stats.GetSanity()},
		{"physical", stats.GetPhysical()},
		{"wit", stats.GetWit()},
		{"composure", stats.GetComposure()},
		{"social", stats.GetSocial()},
	}

	for _, field := range statFields {
		stmt := `INSERT INTO character_stat (game_id, character_id, characteristic, value) VALUES ($1, $2, $3, $4)`
		_, err := pool.Exec(context.Background(), stmt, req.GameID, req.CharacterId, field.name, field.value)
		if err != nil {
			log.Fatal(err)
		}
	}
	res := &pb.CreateCharacterResponse{
		CharacterId: req.CharacterId,
	}
	return res, nil
}

func (s *server) GetAdventureLore(ctx context.Context, req *pb.GetAdventureLoreRequest) (*pb.GetAdventureLoreResponse, error) {
	sql := query := `
		SELECT content
		FROM (
			SELECT lore_content_id, content, embedding,
				   embedding <#> \$2::vector AS cosine_similarity
			FROM lore_content
			WHERE lore_id = \$1
		) AS subquery
		ORDER BY cosine_similarity DESC
		LIMIT 1;`
	row := pool.QueryRow(context.Background(), sql, rqd.LoreID, req.Embedding)
	var content string
	err = row.Scan(&content)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("No matching lore content found")
		} else {
			log.Fatal(err)
		}
	res := &pb.GetAdventureLoreResponse{
		LoreId:  req.LoreId,
		Context: content
	}
	return res, nil
}

func (s *server) GetConversationHistory(ctx context.Context, req *pb.GetConversationHistoryRequest) (*pb.GetConversationHistoryResponse, error) {
	sql := `
		SELECT role, content
		FROM narrative
		WHERE game_id = $1
		ORDER BY created_at ASC;
	`

	rows, err := s.Pool.Query(ctx, sql, req.GameId)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	history := []*pb.Chat{}
	for rows.Next() {
		var role string
		var content string

		if err := rows.Scan(&role, &content); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		chat := &pb.Chat{
			Content: content,
			Role:    pb.Role(pb.Role_value[role]),
		}

		history = append(history, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	res := &pb.GetConversationHistoryResponse{
		GameId: req.GameId,
		History: history,
	}

	return res, nil
}

func (s *server) GetPlayerCharacteristics(ctx context.Context, req *pb.GetPlayerCharacteristicsRequest) (*pb.GetPlayerCharacteristicsResponse, error) {
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
}

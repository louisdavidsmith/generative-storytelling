package server

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	pb "query/pkg/proto/query"
)

type Server struct {
	pb.UnimplementedQueryServiceServer
	Pool *pgxpool.Pool
}

func (s *Server) CreateAdventure(ctx context.Context, req *pb.CreateAdventureRequest) (*pb.CreateAdventureResponse, error) {
	sql := `INSERT INTO adventure (adventure_id, lore_id, premise, adventure_name) VALUES ($1, $2, $3, $4)`
	_, err := s.Pool.Exec(ctx, sql, req.AdventureId, req.LoreId, req.Premise, req.AdventureName)
	if err != nil {
		log.Fatal(err)
	}

	res := &pb.CreateAdventureResponse{
		AdventureId: req.AdventureName,
	}
	return res, nil
}

func (s *Server) CreateAdventureLore(ctx context.Context, req *pb.CreateAdventureLoreRequest) (*pb.CreateAdventureLoreResponse, error) {
	sql := `INSERT INTO lore (lore_id, lore_name, lore_description) VALUES ($1, $2, $3, $4)`
	_, err := s.Pool.Exec(ctx, sql, req.LoreId, req.LoreName, req.LoreDescription)
	if err != nil {
		log.Fatal(err)
	}
	res := &pb.CreateAdventureLoreResponse{
		LoreName: req.LoreId,
	}
	return res, nil
}

func (s *Server) GetAdventureLore(ctx context.Context, req *pb.GetAdventureLoreRequest) (*pb.GetAdventureLoreResponse, error) {
	sql := `
		SELECT content
		FROM (
			SELECT lore_content_id, content, embedding,
				   embedding <#> $2::vector AS cosine_similarity
			FROM lore_content
			WHERE lore_id = $1
		) AS subquery
		ORDER BY cosine_similarity DESC
		LIMIT 1;`
	row := s.Pool.QueryRow(ctx, sql, req.LoreId, req.Embedding)
	var content string
	_ = row.Scan(&content)
	res := &pb.GetAdventureLoreResponse{
		LoreId:  req.LoreId,
		Context: content,
	}
	return res, nil
}

func (s *Server) GetConversationHistory(ctx context.Context, req *pb.GetConversationHistoryRequest) (*pb.GetConversationHistoryResponse, error) {
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
		GameId:  req.GameId,
		History: history,
	}

	return res, nil
}

func (s *Server) GetPlayerCharacteristics(ctx context.Context, req *pb.GetPlayerCharacteristicsRequest) (*pb.GetPlayerCharacteristicsResponse, error) {
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

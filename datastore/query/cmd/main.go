package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	pb "query/pkg/proto/query"
	"query/pkg/server" // Import the server package
)

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
	queryServer := &server.Server{Pool: dbpool} // Create an instance of your server
	pb.RegisterQueryServiceServer(s, queryServer) // Register the server with the gRPC server

	fmt.Println("Server is running on port 50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

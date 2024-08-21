package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    pb "your_project/pkg/proto" // Adjust this import according to your setup
    "your_project/internal/service"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterPromptServiceServer(s, &service.Server{
        SharedData: make(map[string]string), // Initialize shared data
    })

    log.Printf("Server is running on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

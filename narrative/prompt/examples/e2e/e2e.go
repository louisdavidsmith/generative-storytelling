package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "prompt/pkg/proto/prompt" // Adjust this path according to your setup
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPromptServiceClient(conn)

	// Create a request
	req := &pb.GetNarrationPromptRequest{
		GameState:   "Your game state information here",
		LoreContext: "Your mythos setting information here",
		UserInput:   "Your user input here",
		ChatHistory: []*pb.Chat{}, // Provide chat history if needed
	}

	// Set a deadline to the context, to avoid waiting indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the GetNarrationPrompt method
	resp, err := client.GetNarrationPrompt(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get narration prompt: %v", err)
	}

	// Print the response
	for _, chat := range resp.Prompt {
		fmt.Printf("Role: %s\nContent: %s\n\n", chat.Role, chat.Content)
	}
}

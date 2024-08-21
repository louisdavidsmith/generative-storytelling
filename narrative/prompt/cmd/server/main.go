package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "prompt/pkg/proto/prompt" // Adjust this path according to your setup
)

// server is used to implement pb.PromptServiceServer.
type server struct {
	pb.UnimplementedPromptServiceServer
}

func (s *server) GetNarrationPrompt(ctx context.Context, req *pb.GetNarrationPromptRequest) (*pb.Prompt, error) {
	systemPrompt := `You are a subsystem for a lovecraftian roleplaying game of dark and foreboding adventure set in the 1920s.
	Your task is to help players explore scenarios in which they confront and attempt to overcome cosmic evils beyond human comprehension.
	Your role is to seamlessly blend information retrieval and creative generation to enhance the storytelling experience.
	Do not regurgitate information from the mythos, instead use the mythos to creatively inform novel elements of your storytelling.
	When responding to prompts, only use the provided mythos information if it makes sense given the context of the story and the user input.
	After a user has prompted their intentions, you will set out to resolve actions, prompting the user for further input as necessary to
	empower them to tell their own story. When responding you will write no more than two paragraphs at a time,
	finishing each generation by outputing "What do you want to do next?" Additionally, when responding avoid any notes, comments,
	or asides that might take the player out of the story. Here is a structured dictionary that gives information to work into your story. %s`

	formattedSystemPrompt := fmt.Sprintf(systemPrompt, req.GameState)

	systemChat := &pb.Chat{
		Content: formattedSystemPrompt,
		Role:    pb.Role_SYSTEM,
	}

	var last10ChatHistory []*pb.Chat
	if len(req.ChatHistory) > 10 {
		last10ChatHistory = req.ChatHistory[len(req.ChatHistory)-10:]
	} else {
		last10ChatHistory = req.ChatHistory
	}

	userPrompt := `MYTHOS_SETTING: %s | USER_INPUT %s`
	formattedUserPrompt := fmt.Sprintf(userPrompt, req.LoreContext, req.UserInput)
	userChat := &pb.Chat{
		Content: formattedUserPrompt,
		Role:    pb.Role_USER,
	}
	prompt := &pb.Prompt{
		Prompt: append([]*pb.Chat{systemChat}, append(last10ChatHistory, userChat)...),
	}

	return prompt, nil
}

func (s *server) GetSkillCheckResolutionPrompt(ctx context.Context, req *pb.GetSkillCheckResolutionPromptRequest) (*pb.Prompt, error) {
	// Implement the logic for handling the GetSkillCheckResolutionPrompt request.
	// For demonstration, let's just create a simple response.
	prompt := &pb.Prompt{
		Prompt: []*pb.Chat{
			{
				Content: "This is a skill check resolution prompt based on the check outcome.",
				Role:    pb.Role_ASSISTANT,
			},
		},
	}

	return prompt, nil
}

func (s *server) GetInsanityNarration(ctx context.Context, req *pb.GetSanityNarrationRequest) (*pb.Prompt, error) {
	// Implement the logic for handling the GetInsanityNarration request.
	// For demonstration, let's just create a simple response.
	prompt := &pb.Prompt{
		Prompt: []*pb.Chat{
			{
				Content: "This is an insanity narration prompt based on the current sanity level.",
				Role:    pb.Role_ASSISTANT,
			},
		},
	}

	return prompt, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPromptServiceServer(s, &server{})

	log.Printf("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

package service

import (
    "context"
    "sync"

    pb "your_project/pkg/proto" // Adjust this import according to your setup
)

// Server is used to implement pb.PromptServiceServer.
type Server struct {
    pb.UnimplementedPromptServiceServer
    SharedData map[string]string // Example shared data
}

func (s *Server) GetNarrationPrompt(ctx context.Context, req *pb.GetNarrationPromptRequest) (*pb.Prompt, error) {

    // Access or modify shared data if needed
    prompt := &pb.Prompt{
        Prompt: []*pb.Chat{
            {
                Content: "This is a narration prompt based on the provided lore context and game state.",
                Role:    pb.Role_ASSISTANT,
            },
        },
    }

    return prompt, nil
}

func (s *Server) GetSkillCheckResolutionPrompt(ctx context.Context, req *pb.GetSkillCheckResolutionPromptRequest) (*pb.Prompt, error) {

    // Access or modify shared data if needed
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

func (s *Server) GetInsanityNarration(ctx context.Context, req *pb.GetSanityNarrationRequest) (*pb.Prompt, error) {

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

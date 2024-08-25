package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"

	pb "prompt/pkg/proto/prompt"
)

type Config struct {
	NarrativePromptPort          int
	MaxConversationHistoryTokens int
}

type server struct {
	pb.UnimplementedPromptServiceServer
	Config *Config
}

func LoadConfig() (*Config, error) {
	PromptPortStr := os.Getenv("NARRATIVE_PROMPT_PORT")
	PromptPort, err := strconv.Atoi(PromptPortStr)

	if err != nil {
		return nil, fmt.Errorf("NARRATIVE_PROMPT_PORT environment variable not set or not an integer")
	}

	MaxTokensStr := os.Getenv("MAX_CONVESATION_HISTORY_TOKENS")
	MaxTokens, err := strconv.Atoi(MaxTokensStr)
	if err != nil {
		return nil, fmt.Errorf("MAX_CONVESATION_HISTORY_TOKENS environment variable not set or not an integer")
	}
	config := &Config{
		NarrativePromptPort:          PromptPort,
		MaxConversationHistoryTokens: MaxTokens,
	}

	return config, nil
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

	userPrompt := `MYTHOS_SETTING: %s | USER_INPUT %s`
	formattedUserPrompt := fmt.Sprintf(userPrompt, req.LoreContext, req.UserInput)
	userChat := &pb.Chat{
		Content: formattedUserPrompt,
		Role:    pb.Role_USER,
	}
	prompt := &pb.Prompt{
		Prompt: append([]*pb.Chat{systemChat}, append(req.ChatHistory, userChat)...),
	}

	return prompt, nil
}

func (s *server) GetSkillCheckResolutionPrompt(ctx context.Context, req *pb.GetSkillCheckResolutionPromptRequest) (*pb.Prompt, error) {
	skillCheckPrompt := `You are a subsystem for a lovecraftian
                   role-playing game. When the player is exploring the world,
                   they will at times take actions that must be resolved via a
                   skill check. Your purpose is to take the player action and
                   the result of the skill check and provide a short narration
                   to resolve the action according to the skill check.
		   Check Result %s | User Action %s`
	formattedSkillPrompt := fmt.Sprintf(skillCheckPrompt, req.CheckOutcome, req.UserInput)
	skillChat := &pb.Chat{
		Content: formattedSkillPrompt,
		Role:    pb.Role_USER,
	}

	prompt := &pb.Prompt{
		Prompt: append(req.ChatHistory, skillChat),
	}
	return prompt, nil
}

func (s *server) GetInsanityNarration(ctx context.Context, req *pb.GetSanityNarrationRequest) (*pb.Prompt, error) {
	sanityPrompt := `You are a subsystem for a lovecraftian role-playing game.
		When a player encounters horrifying and eldritch powers beyond human
		comprehension, they must undergo a sanity check. If a player fails
		that check, your subsystem will take the current context of the adventure,
		setting information from the lovecraftian mythos, and write a scene that
		describes a temporary and transient bout of madness. When completing a request
		modulate the intensity of your response based on the severity of the situation,
		and the nature of the threat. When responding do not refer to yourself or have any
		asides, merely continue the narrative with a bout of madness. MYTHOS: %s | USER INPUT %s`

	formattedSanityPrompt := fmt.Sprintf(sanityPrompt, req.LoreContent, req.UserInput)
	sanityChat := &pb.Chat{
		Content: formattedSanityPrompt,
		Role:    pb.Role_USER,
	}

	prompt := &pb.Prompt{
		Prompt: append(req.ChatHistory, sanityChat),
	}

	return prompt, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	config, _ := LoadConfig()
	pb.RegisterPromptServiceServer(s, &server{Config: config})

	log.Printf("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

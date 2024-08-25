package mistral

import (
	"context"
	pb "generation/pkg/proto/generation"
	"github.com/gage-technologies/mistral-go"
	"log"
)

type MistralGenerator struct {
	Client  *mistral.MistralClient
	ModelId string
}

func convertChatToChatMessage(chat *pb.Chat) mistral.ChatMessage {
	role := ""
	switch chat.Role {
	case pb.Role_USER:
		role = "USER"
	case pb.Role_ASSISTANT:
		role = "ASSISTANT"
	case pb.Role_SYSTEM:
		role = "SYSTEM"
	default:
		log.Fatalf("Unknown role: %d", chat.Role)
	}

	return mistral.ChatMessage{ // Remove the pointer here
		Role:    role,
		Content: chat.Content,
	}
}

func convertChatsToChatMessages(chatList []*pb.Chat) []mistral.ChatMessage {
	chatMessages := make([]mistral.ChatMessage, len(chatList)) // Change the type of the slice
	for i, chat := range chatList {
		chatMessages[i] = convertChatToChatMessage(chat)
	}
	return chatMessages
}

func (m *MistralGenerator) GenerateResponseStream(req *pb.GenerateResponseRequest, stream pb.GenerationService_GenerateResponseStreamServer) error {
	log.Printf("Received request: %v", req)
	chatInput := convertChatsToChatMessages(req.Chat)
	chatResChan, err := m.Client.ChatStream(m.ModelId, chatInput, nil)
	if err != nil {
		log.Fatalf("Error getting chat completion stream: %v", err)
	}
	for chatResChunk := range chatResChan {
		if chatResChunk.Error != nil {
			log.Fatalf("Error while streaming response: %v", chatResChunk.Error)
		}
		stream.Send(&pb.StreamResponse{Token: chatResChunk.Choices[0].Delta.Content})
	}
	return nil
}

func (m *MistralGenerator) GenerateResponseBatch(ctx context.Context, req *pb.GenerateResponseRequest) (*pb.BatchResponse, error) {
	log.Printf("Received request: %v", req)
	chatInput := convertChatsToChatMessages(req.Chat)
	chatRes, err := m.Client.Chat(m.ModelId, chatInput, nil)
	if err != nil {
		log.Fatalf("Error getting chat completion: %v", err)
	}
	return &pb.BatchResponse{Response: chatRes.Choices[0].Message.Content}, nil
}

func NewMistralGenerator(client *mistral.MistralClient, modelId string) *MistralGenerator {
	return &MistralGenerator{
		Client:  client,
		ModelId: modelId,
	}
}

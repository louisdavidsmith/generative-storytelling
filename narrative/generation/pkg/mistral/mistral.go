package mistral

import (
	"context"
	pb "generation/pkg/proto/generation"
	"github.com/gage-technologies/mistral-go"
	"log"
)

type MistralGenerator struct {
	Client  *mistral.NewMistralClientDefault
	ModelId string
}

func convertChatToChatMessage(chat *pb.Chat) *ChatMessage {
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

	return &ChatMessage{
		Role:    role,
		Content: chat.Content,
	}
}

func convertChatsToChatMessages(chatList []*pb.Chat) []*ChatMessage {
	chatMessages := make([]*ChatMessage, len(chatList))
	for i, chat := range chatList {
		chatMessages[i] = convertChatToChatMessage(chat)
	}
	return chatMessages
}

func (m *MistralGenerator) GenerateResponseStream(req *pb.GenerateResponseRequest, stream pb.GenerationService_GenerateResponseStreamServer) error {
	log.Printf("Received request: %v", req)
	chatInput = convertChatsToChatMessages(req.Chat)
	chatResChan, err := client.ChatStream(m.ModelId, chatInput, nil)
	if err != nil {
		log.Fatalf("Error getting chat completion stream: %v", err)
	}
	for chatResChunk := range chatResChan {
		if chatResChunk.Error != nil {
			log.Fatalf("Error while streaming response: %v", chatResChunk.Error)
		}
		stream.Send(&pb.StreamResponse{Token: chatResChunk})

	}
	return nil
}

func (m *MistralGenerator) GenerateResponseBatch(ctx context.Context, req *pb.GenerateResponseRequest) (*pb.BatchResponse, error) {
	log.Printf("Received request: %v", req)
	chatInput = convertChatsToChatMessages(req.Chat)
	chatRes, err := client.Chat(m.ModelId, chatInput, nil)
	if err != nil {
		log.Fatalf("Error getting chat completion: %v", err)
	}
	return &pb.BatchResponse{Response: chatRes}, nil
}

func NewMistralGenerator(client *mistral.NewMistralClientDefault, modelId string) *MistralGenerator {
	return &MistralGenerator{
		Client:  client,
		ModelId: modelId,
	}
}

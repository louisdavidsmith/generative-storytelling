package truncate

import (
	pb "prompt/pkg/proto/prompt"
)

func TruncateConversationHistory(history []*pb.Chat, maxLength int) []*pb.Chat {
	// a token is roughly four characters
	maxChars := maxLength * 4

	var truncatedHistory []*pb.Chat
	totalChars := 0

	for i := len(history) - 1; i >= 0; i-- {
		chat := history[i]
		chatLength := len(chat.Content)
		if totalChars+chatLength <= maxChars {
			totalChars += chatLength
			truncatedHistory = append(truncatedHistory, chat)
		}
	}

	for i, j := 0, len(truncatedHistory)-1; i < j; i, j = i+1, j-1 {
		truncatedHistory[i], truncatedHistory[j] = truncatedHistory[j], truncatedHistory[i]
	}

	return truncatedHistory
}

package interfaces

import (
	"skripsi/features/chatbot/entity"

	"github.com/sashabaranov/go-openai"
)

type ChatbotServiceInterface interface {
	HandleCustomerQuery(query string) (string, []string, error)
	GetCompletionFromMessages(request entity.ChatRequest) (openai.ChatCompletionResponse, error)
}

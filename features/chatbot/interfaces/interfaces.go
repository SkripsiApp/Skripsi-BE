package interfaces

import (
	"skripsi/features/chatbot/entity"

	"github.com/sashabaranov/go-openai"
)

type ChatbotServiceInterface interface {
	HandleCustomerQuery(query string) (string, []entity.ProductRecommendation, error)
	GetCompletionFromMessages(request entity.ChatRequest) (openai.ChatCompletionResponse, error)
}

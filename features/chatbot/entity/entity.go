package entity

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type ChatRequest struct {
	Context  context.Context
	Client   *openai.Client
	Messages []openai.ChatCompletionMessage
	Model    string
}

type ProductRecommendation struct {
	Name  string
	Price int
	Image string
}
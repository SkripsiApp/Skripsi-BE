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
	Id    string
	Name  string
	Price int
	Sold  int
	Image string
}

package request

type ChatbotRequest struct {
	Question string `json:"question" validate:"required"`
}

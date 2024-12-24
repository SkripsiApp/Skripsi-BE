package controller

import (
	"net/http"
	"skripsi/features/chatbot/dto/request"
	"skripsi/features/chatbot/interfaces"

	"github.com/labstack/echo/v4"
)

type chatbotController struct {
	chatbotService interfaces.ChatbotServiceInterface
}

func NewChatbotController(chatbotService interfaces.ChatbotServiceInterface) *chatbotController {
	return &chatbotController{
		chatbotService: chatbotService,
	}
}

func (c *chatbotController) HandleQuery(e echo.Context) error {
	var req request.ChatbotRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if req.Question == "" {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "Question is required"})
	}

	answer, imageURLs, err := c.chatbotService.HandleCustomerQuery(req.Question)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := map[string]interface{}{
		"answer": answer,
		"images": imageURLs,
	}

	return e.JSON(http.StatusOK, response)
}

package controller

import (
	"net/http"
	"skripsi/features/chatbot/dto/request"
	"skripsi/features/chatbot/dto/response"
	"skripsi/features/chatbot/interfaces"
	"skripsi/utils/helper"

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
		return helper.ResponseError(400, "Invalid input data")
	}

	if req.Question == "" {
		return helper.ResponseError(400, "Pertanyaan tidak boleh kosong")
	}

	answer, imageURLs, err := c.chatbotService.HandleCustomerQuery(req.Question)
	if err != nil {
		return helper.ResponseError(500, err.Error())
	}

	response := response.ChatbotResponse{
		Answer: answer,
		Images: imageURLs,
	}

	return e.JSON(http.StatusOK, helper.ResponseSuccessWithData("berhasil mendapatkan response", response))
}

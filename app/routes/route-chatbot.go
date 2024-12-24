package routes

import (
	"os"
	"skripsi/features/chatbot/controller"
	"skripsi/features/chatbot/service"
	"skripsi/features/product/repository"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteChatbot(e *echo.Group, db *gorm.DB) {
	productRepository := repository.NewProductRepository(db)

	godotenv.Load()
	openaiKey := os.Getenv("OPENAI_API_KEY")

	chatbotService := service.NewChatbotService(openaiKey, productRepository)
	chatbotController := controller.NewChatbotController(chatbotService)

	chatbot := e.Group("/chatbot")
	chatbot.POST("", chatbotController.HandleQuery)
}

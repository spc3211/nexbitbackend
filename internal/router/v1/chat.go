package v1

import (
	handler "nexbit/internal/handler/requestHandler"
	chatService "nexbit/internal/service"

	"github.com/gofiber/fiber/v2"
)

func ChatRouter(app *fiber.App, chatService *chatService.ChatService) {

	handler := handler.NewChatHandler(chatService)
	api := app.Group("/v1")
	api.Post("/chat/chat-complete", handler.ChatHandler)
	api.Post("/stock/get-fundamentals", handler.FundamentalHandler)
	api.Get("/stock/news-insights", handler.NewsInsightsHandler)
}

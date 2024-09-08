package v1

import (
	requesthandler "nexbit/internal/handler/requesthandler"
	chatService "nexbit/internal/service"

	"github.com/gofiber/fiber/v2"
)

func ChatRouter(app *fiber.App, chatService *chatService.ChatService) {

	handler := requesthandler.NewChatHandler(chatService)
	api := app.Group("/v1")
	api.Post("/chat/chat-complete", handler.UserChatHandler)
	api.Get("/stock/get-fundamentals", handler.FundamentalHandler)
}

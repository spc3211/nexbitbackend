package v1

import (
	requesthandler "nexbit/internal/handler/requesthandler"
	chatService "nexbit/internal/service"

	"net/http"
	"io"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ChatRouter(app *fiber.App, chatService *chatService.ChatService) {

	handler := requesthandler.NewChatHandler(chatService)
	api := app.Group("/v1")
	api.Get("/stock/news-insights", handler.NewsInsightsHandler)
	api.Post("/chat/chat-complete", handler.UserChatHandler)
	api.Get("/stock/get-fundamentals", handler.FundamentalHandler)
	api.Post("/stock/save-reports", handler.FileUploadHandler)
	api.Post("/chat/ask", handler.UserQueryHandler)

	api.Get("/debug", func(c *fiber.Ctx) error {
		url := "https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers=TSLA&limit=5&time_from=20240915T2011&apikey=JF509SILZGIN3B5O"
		resp, err := http.Get(url)
		if err != nil {
			// Respond with error message
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error: %v", err))
		}
		defer resp.Body.Close()

		// Read the response body
		body, _ := io.ReadAll(resp.Body)

		// Respond with the API response
		return c.SendString(string(body))
	})
	
}

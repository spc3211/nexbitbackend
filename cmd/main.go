package main

import (
	"fmt"
	"time"

	router "nexbit/internal/router/v1"
	chatService "nexbit/internal/service"

	logger "nexbit/util"

	openai "github.com/sashabaranov/go-openai"

	externalFmpApiClient "nexbit/external/fmp"
	externalOpenAiClient "nexbit/external/openai"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	fmt.Println("Starting the server...")
	logger.Init()
	app := fiber.New()

	// Add the recover middleware
	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: stackTraceHandler,
	}))

	openaiClient := openai.NewClient("s")

	externalChatGptClient := externalOpenAiClient.NewOpenAiClient(openaiClient)
	httpClient := externalFmpApiClient.NewHTTPClient(5 * time.Second)
	externalFmpApiClient := externalFmpApiClient.NewAPIClient(httpClient)

	chatService := chatService.NewChatService(externalChatGptClient, externalFmpApiClient)
	router.ChatRouter(app, chatService)

	if err := app.Listen(":3002"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func stackTraceHandler(ctx *fiber.Ctx, err interface{}) {
	errMsg := fmt.Sprintf("Panic: %v", err)
	ctx.Status(fiber.StatusInternalServerError)
	err = ctx.JSON(fiber.Map{
		"error":   errMsg,
		"message": "Internal Server Error. Please try again later.",
	})
}

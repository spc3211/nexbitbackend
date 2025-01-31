package main

import (
	"fmt"
	"os"
	"time"

	"nexbit/internal/repo"
	router "nexbit/internal/router/v1"
	service "nexbit/internal/service"

	logger "nexbit/util"

	external "nexbit/external"
	externalFmpApiClient "nexbit/external/fmp"
	externalNewsClient "nexbit/external/news"
	externalOpenAiClient "nexbit/external/openai"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"

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

	//connect database
	//connStr := "user=nexbit dbname=chat password=password host=localhost port=5432 sslmode=disable"

	// Loads the .env file
	_ = godotenv.Load()

	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	openAiApiKey := os.Getenv("OPENAI_API_KEY")

	if port == "" {
		port = "3001"
	}

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
		user, dbname, password, host, port, sslmode)

	dbService, err := repo.NewDBService(connStr)
	if err != nil {
		fmt.Println("error in db service")
	}

	// err = dbService.Ping()
	// if err != nil {
	// 	log.Fatalf("Error connecting to the database: %v\n", err)
	// } else {
	// 	fmt.Println("Successfully connected to the PostgreSQL database!")
	// }

	defer dbService.Close()

	openaiClient := openai.NewClient(openAiApiKey)

	externalChatGptClient := externalOpenAiClient.NewOpenAiClient(openaiClient)
	httpClient := external.NewHTTPClient(5 * time.Second)
	externalFmpApiClient := externalFmpApiClient.NewAPIClient(httpClient)
	externalNewsApiClient := externalNewsClient.NewAPIClient(httpClient)

	chatService := service.NewChatService(dbService, externalChatGptClient, externalFmpApiClient, externalNewsApiClient)
	onboardingService := service.NewOnboardingService(dbService)
	router.ChatRouter(app, chatService)
	router.OnboardingRouter(app, onboardingService)

	if err := app.Listen(":" + port); err != nil {
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

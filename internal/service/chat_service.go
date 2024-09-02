package service

import (
	"fmt"
	openAiClient "nexbit/external/openai"
	"nexbit/util"

	models "nexbit/models"

	"github.com/gofiber/fiber/v2"
)

type ChatService struct {
	openAiClient *openAiClient.OpenAiClient
}

func NewChatService(openAiClient *openAiClient.OpenAiClient) *ChatService {
	return &ChatService{openAiClient: openAiClient}
}

func (s *ChatService) ChatService(ctx *fiber.Ctx) error {
	message := ctx.Locals("requestData").(models.SubmitChatRequest)

	chatResponse, err := s.openAiClient.ChatCompletionClient(ctx.Context(), message.Message)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[ChatService] Failed to process chat request. err: %v", err)
		return err
	}
	fmt.Println(chatResponse.Choices[0].Message)

	return ctx.Context().Err()
}

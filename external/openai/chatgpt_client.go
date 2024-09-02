package ory

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAiClient struct {
	apiClient *openai.Client
}

func NewOpenAiClient(apiClient *openai.Client) *OpenAiClient {
	return &OpenAiClient{
		apiClient: apiClient,
	}
}

func (h *OpenAiClient) ChatCompletionClient(ctx context.Context, message string) (openai.ChatCompletionResponse, error) {
	resp, err := h.apiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return resp, fmt.Errorf("[ChatCompletionClient] Error while calling chatgpt : %v", err)
	}

	return resp, nil

}

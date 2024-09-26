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

func (h *OpenAiClient) ChatCompletionClient(ctx context.Context, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	resp, err := h.apiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4oLatest,
			Messages: messages,
		},
	)

	if err != nil {
		return resp, fmt.Errorf("[ChatCompletionClient] Error while calling chatgpt : %v", err)
	}

	return resp, nil
}

func (h *OpenAiClient) UploadFileClient(ctx context.Context, req openai.FileRequest) (openai.File, error) {
	fileResp, err := h.apiClient.CreateFile(context.Background(), req)

	if err != nil {
		return fileResp, fmt.Errorf("[UploadFileClient] error uploading file: %w", err)
	}
	return fileResp, nil
}

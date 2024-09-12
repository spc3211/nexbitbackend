package service

import (
	"fmt"
	fmpApiClient "nexbit/external/fmp"
	openAiClient "nexbit/external/openai"
	newsApiClient "nexbit/external/news"
	"nexbit/util"

	models "nexbit/models"

	"github.com/gofiber/fiber/v2"
)

type ChatService struct {
	openAiClient     *openAiClient.OpenAiClient
	fmpApiClient     *fmpApiClient.FmpApiClient
	newsApiClient    *newsApiClient.NewsApiClient
}

func NewChatService(openAiClient *openAiClient.OpenAiClient, fmpApiClient *fmpApiClient.FmpApiClient, newsApiClient *newsApiClient.NewsApiClient) *ChatService {
	return &ChatService{
		openAiClient:     openAiClient,
		fmpApiClient:     fmpApiClient,
		newsApiClient:    newsApiClient,
	}
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

func (s *ChatService) FetchFundamentals(ctx *fiber.Ctx) error {
	stockSymbol := ctx.Locals("stockSymbol").(string)

	incomeStatementResponse, err := s.fmpApiClient.FetchIncomeStatementAPI(ctx.Context(), stockSymbol, "annual")
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[ChatService] Failed to process chat request. err: %v", err)
		return err
	}

	balanceSheetResponse, err := s.fmpApiClient.FetchBalanceSheet(ctx.Context(), stockSymbol, "annual")
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[ChatService] Failed to process chat request. err: %v", err)
		return err
	}

	cashFlowResponse, err := s.fmpApiClient.FetchCashFlowStatement(ctx.Context(), stockSymbol, "annual")
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[ChatService] Failed to process chat request. err: %v", err)
		return err
	}

	finalRespnse := models.FundamentalDataResponse{
		BalanceSheetResponse:    balanceSheetResponse,
		IncomeStatementResponse: incomeStatementResponse,
		CashFlowResponse:        cashFlowResponse,
	}

	return ctx.JSON(fiber.Map{
		"stock_financials": finalRespnse,
	})
}

func (s *ChatService) FetchNewsInsights(ctx *fiber.Ctx) error {
	stockSymbol := ctx.Locals("stockSymbol").(string)

	insights, err := s.newsApiClient.FetchNewsInsights(ctx.Context(), stockSymbol)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[ChatService] Failed to process chat request. err: %v", err)
		return err
	}

	return ctx.JSON(insights)
}

package service

import (
	"context"
	"encoding/json"
	"fmt"
	fmpApiClient "nexbit/external/fmp"
	newsApiClient "nexbit/external/news"
	openAiClient "nexbit/external/openai"
	"nexbit/internal/repo"
	"nexbit/util"
	"path/filepath"
	"time"

	models "nexbit/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sashabaranov/go-openai"
)

type ChatService struct {
	openAiClient  *openAiClient.OpenAiClient
	fmpApiClient  *fmpApiClient.FmpApiClient
	newsApiClient *newsApiClient.NewsApiClient
	db            *repo.DBService
}

func NewChatService(db *repo.DBService, openAiClient *openAiClient.OpenAiClient, fmpApiClient *fmpApiClient.FmpApiClient, newsApiClient *newsApiClient.NewsApiClient) *ChatService {
	return &ChatService{
		openAiClient:  openAiClient,
		fmpApiClient:  fmpApiClient,
		newsApiClient: newsApiClient,
		db:            db,
	}
}

func (s *ChatService) ChatService(ctx *fiber.Ctx) error {
	messages := ctx.Locals("requestData").(models.SubmitChatRequest)
	var chatgptMessages []openai.ChatCompletionMessage

	for _, message := range messages.Message {
		var chatgptMessage openai.ChatCompletionMessage
		chatgptMessage.Role = message.Role
		chatgptMessage.Content = message.Content
		chatgptMessages = append(chatgptMessages, chatgptMessage)
	}

	chatResponse, err := s.openAiClient.ChatCompletionClient(ctx.Context(), chatgptMessages)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[ChatService] Failed to process chat request. err: %v", err)
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": chatResponse.Choices[0].Message,
	})
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

	financialRationResponse, err := s.fmpApiClient.FetchFinancialsRatio(ctx.Context(), stockSymbol, "annual")
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[ChatService] Failed to process chat request. err: %v", err)
		return err
	}

	finalRespnse := models.FundamentalDataResponse{
		BalanceSheetResponse:    balanceSheetResponse,
		IncomeStatementResponse: incomeStatementResponse,
		CashFlowResponse:        cashFlowResponse,
		FinancialRatiosResponse: financialRationResponse,
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

func (s *ChatService) Uploadfile(ctx *fiber.Ctx, req models.FileUploadRequest) error {
	fileInfo, err := getFileInfos(req.FilePaths)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[Uploadfile] Failed to get file infos. err: %v", err)
		return err
	}

	var fileIDs []string
	for _, info := range fileInfo {
		fileReq := openai.FileRequest{
			FileName: info.Name,
			FilePath: info.Path,
			Purpose:  "assistants",
		}

		file, err := s.openAiClient.UploadFileClient(ctx.Context(), fileReq)
		if err != nil {
			util.WithContext(ctx.Context()).Errorf("[Uploadfile] Failed to upload file: %s. err: %v", info.Name, err)
			return err
		}
		fileIDs = append(fileIDs, file.ID)
	}

	time.Sleep(10 * time.Second)

	// Prepare the chat message with content generated by the separate method
	chatgptMessages := []openai.ChatCompletionMessage{
		{
			Role:    "user",
			Content: s.generatePromptContent(fileIDs),
		},
	}

	chatResponse, err := s.openAiClient.ChatCompletionClient(ctx.Context(), chatgptMessages)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[Uploadfile] Failed to process chat request. err: %v", err)
		return err
	}

	fmt.Println(chatResponse.Choices[0].Message.Content)

	var response models.StockResearchResponse
	if err := json.Unmarshal([]byte(chatResponse.Choices[0].Message.Content), &response); err != nil {
		return fmt.Errorf("[Uploadfile] Error parsing JSON: %v", err)
	}

	if response.Err != "" {
		return fmt.Errorf("[Uploadfile] Tailored error response from gpt: %v", err)
	}

	for _, reportResponse := range response.Data {

		dbReq := repo.StockResearchReport{
			Company:            reportResponse.Company,
			Sector:             reportResponse.Sector,
			Recommendation:     reportResponse.Recommendation,
			TargetPrice:        reportResponse.TargetPrice,
			RevenueProjections: reportResponse.RevenueProjections,
			CAGR:               reportResponse.CAGR,
			EBITDA:             reportResponse.EBITDA,
			NewsSummary:        reportResponse.NewsSummary,
		}

		err := s.db.SaveStockReport(context.Background(), dbReq)
		if err != nil {
			util.WithContext(ctx.Context()).Errorf("[Uploadfile] Failed to save data in database. err: %v", err)
			return err
		}

	}

	return nil
}

func (s *ChatService) generatePromptContent(fileIDs []string) string {
	return fmt.Sprintf("You have been provided with a list of file IDs corresponding to stock research reports. Each research report contains detailed information about the performance of Indian publicly listed companies. Your task is to analyze each file thoroughly and extract the relevant data to populate a StockReport object. Please adhere strictly to the following instructions:\n\n1. Extract only verified and factual information directly from the reports—do not infer or assume any details.\n2. For each report, fill out the following StockReport struct:\n   \n    type StockReport struct {\n        Company            string    `json:\"company\"`           // Name of the company (usually in the title or company info section)\n        Sector             string    `json:\"sector\"`            // Industry sector (found near the company name or at the top of the report)\n        Recommendation     string    `json:\"recommendation\"`    // Analyst recommendation (e.g., Buy, Hold, Sell, often in the first few pages)\n        TargetPrice        float64   `json:\"target_price\"`      // Target price in INR (commonly near the recommendation)\n        RevenueProjections []float64 `json:\"revenue_projections\"` // Revenue projections for future years (usually found in financial projections section)\n        CAGR               float64   `json:\"cagr\"`              // Compound Annual Growth Rate (CAGR, found in financial projections)\n        EBITDA             float64   `json:\"ebitda\"`            // Earnings before Interest, Taxes, Depreciation, and Amortization (financial section)\n        NewsSummary        string    `json:\"news_summary\"`      // Summary of key news related to the company (usually in a company update or news section)\n    }\n    \n3. **Guidance on extraction**:\n    - The **company name** and **sector** are typically found in the title or introductory section of the report.\n    - The **recommendation** and **target price** are located on the first or second page under a headline such as \"BUY,\" \"SELL,\" or \"HOLD.\"\n    - **Revenue projections** are located in the financial tables under \"Net Sales\" or \"Revenue,\" often alongside future fiscal years (FY24, FY25E, FY26E).\n    - **CAGR** and **EBITDA** values are found in the financial projections section, typically as percentages or absolute values.\n    - **News summary** is extracted from the narrative sections of the report, detailing company updates or significant events.\n\n4. Make sure to cross-check each extracted field for accuracy. If data is missing for any field, mark it as `null` or \"N/A\" in the JSON output.\n\n5. The final output must be formatted as JSON, without any additional explanations or assumptions. Here is the required JSON format:\n\n    {\n        \"data\": [ \n            { \"company\": \"string\", \"sector\": \"string\", \"recommendation\": \"string\", \"target_price\": \"float64\", \"revenue_projections\": [ \"float64\" ], \"cagr\": \"float64\", \"ebitda\": \"float64\", \"news_summary\": \"string\" }\n        ],\n        \"err\": null if successful, or a relevant error message if extraction fails.\n    }\n\n6. **IMPORTANT**: Only process the provided file IDs: %v", fileIDs)
}

func getFileInfos(filePaths []string) ([]models.FileInfo, error) {
	var fileInfos []models.FileInfo

	for _, filePath := range filePaths {
		fileName := filepath.Base(filePath)
		fileInfos = append(fileInfos, models.FileInfo{Name: fileName, Path: filePath})
	}

	return fileInfos, nil
}

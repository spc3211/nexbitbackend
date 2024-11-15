package news

import (
	"context"
	"encoding/json"
	"fmt"
	external "nexbit/external"
	models "nexbit/models"
	"time"
)

const API_TOKEN = "C525SCG5OX8NSY1Q"
const BASE_URL = "https://www.alphavantage.co/"

type NewsApiClient struct {
	httpClient *external.HTTPClient
}

func NewAPIClient(httpClient *external.HTTPClient) *NewsApiClient {
	return &NewsApiClient{
		httpClient: httpClient,
	}
}

func (c *NewsApiClient) FetchNewsInsights(ctx context.Context, ticker string) (map[string]interface{}, error) {

	timeTwoMonthsAgo := time.Now().UTC().AddDate(0, -2, 0).Format("20060102T1504")

	url := fmt.Sprintf("%squery?function=NEWS_SENTIMENT&tickers=%s&limit=5&time_from=%s&apikey=%s",
		BASE_URL, ticker, timeTwoMonthsAgo, API_TOKEN)

	fmt.Println("calling news api");
	data, err := c.httpClient.Get(ctx, url, nil)

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	var response models.NewsAPIResponse
	unmarshalErr := json.Unmarshal([]byte(data), &response)
	if unmarshalErr != nil {
		fmt.Println("Error unmarshaling JSON:", unmarshalErr)
		return nil, fmt.Errorf("failed to unmarshal JSON response: %w", unmarshalErr)
	}
	fmt.Println("######## done unmarshalling ########",response);

	var filteredNews []models.FilteredNewsItem
	for _, result := range response.Feed {
		for _, sentiment := range result.TickerSentiment {
			if sentiment.Ticker == ticker {
				filteredNews = append(filteredNews, models.FilteredNewsItem{
					Title:              result.Title,
					TimePublished:      result.TimePublished,
					Sentiment:          sentiment.TickerSentimentScore,
					SentimentRelevance: sentiment.RelevanceScore,
					SentimentLabel:     sentiment.TickerSentimentLabel,
				})
			}
		}
	}

	finalResponse := map[string]interface{}{
		"items":                      len(filteredNews),
		"sentiment_score_definition": "x <= -0.35: Bearish; -0.35 < x <= -0.15: Somewhat-Bearish; -0.15 < x < 0.15: Neutral; 0.15 <= x < 0.35: Somewhat_Bullish; x >= 0.35: Bullish",
		"relevance_score_definition": "0 < x <= 1, with a higher score indicating higher relevance.",
		"feed":                       filteredNews,
	}

	return finalResponse, nil

}

package polygon

import (
	"context"
	"encoding/json"
	"time"
	"fmt"
	models "nexbit/models"
	external "nexbit/external"
)


const API_TOKEN = "6oGHs7F1pbQLMyy5p8i6ST5RJzUEeIAL"
const BASE_URL = "https://api.polygon.io/v2/reference/"

type PolygonApiClient struct {
	httpClient *external.HTTPClient
}

func NewAPIClient(httpClient *external.HTTPClient) *PolygonApiClient {
	return &PolygonApiClient{
		httpClient: httpClient,
	}
}

func (c *PolygonApiClient) FetchNewsInsights(ctx context.Context, ticker string) ([]models.NewsDataInsight, error) {
	url := fmt.Sprintf("%snews?ticker=%s&apiKey=%s&published_utc.gt=%s&order=desc&limit=20&sort=published_utc", BASE_URL,
		ticker, API_TOKEN, time.Now().UTC().AddDate(0, -1, 0).Format(time.RFC3339))

	data, err := c.httpClient.Get(ctx, url, nil)

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	var response models.NewsDataResponse
	unmarshalErr := json.Unmarshal([]byte(data), &response)

	if unmarshalErr != nil {
		fmt.Println("Error unmarshaling JSON:", unmarshalErr)
		return nil, fmt.Errorf("failed to unmarshal JSON response: %w", unmarshalErr)
	}

	var filteredInsights []models.NewsDataInsight
	for _, result := range response.Results {
		for _, insight := range result.Insights {
			if insight.Ticker == ticker {
				filteredInsights = append(filteredInsights, insight)
			}
		}
	}
	return filteredInsights, nil
}
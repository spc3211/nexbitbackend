package repo

type StockResearchReport struct {
	Company            string    `db:"company"`
	Sector             string    `db:"sector"`
	Recommendation     string    `db:"recommendation"`
	TargetPrice        float64   `db:"target_price"`
	RevenueProjections []float64 `db:"revenue_projections"`
	CAGR               float64   `db:"cagr"`
	EBITDA             float64   `db:"ebitda"`
	NewsSummary        string    `db:"news_summary"`
}

type Question struct {
    Question string  `json:"question"`
    Answers  []Answer `json:"answers"`
}

type Answer struct {
    Text  string `json:"text"`
    Score int    `json:"score"`
}

type Portfolio struct {
    Name      string            `json:"name"`
    Allocation map[string]int    `json:"allocation"`
    Sectors   map[string]int    `json:"sectors"`
}

type UserInput struct {
    Answers []int `json:"answers"`
}

type UserPortfolio struct {
    UserID         int    `json:"user_id"`
    PortfolioName  string `json:"portfolio_name"`
}


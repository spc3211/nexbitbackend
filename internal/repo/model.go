package repo

type StockResearchReport struct {
	Id                 string  `db:"id"`
	Company            string  `db:"company"`
	Sector             string  `db:"sector"`
	Recommendation     string  `db:"recommendation"`
	TargetPrice        float64 `db:"target_price"`
	RevenueProjections string  `db:"revenue_projections"`
	CAGR               float64 `db:"cagr"`
	EBITDA             string  `db:"ebitda"`
	Ticker             string  `db:"ticker"`
	Date               string  `json:"date"`
	NewsSummary        string  `db:"news_summary"`
}

type StockResearchFetchRequest struct {
	Date        string
	Sector      string
	CompanyName string
	Ticker      string
}

type Question struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	Text  string `json:"text"`
	Score int    `json:"score"`
}

type Portfolio struct {
	Name       string         `json:"name"`
	Allocation map[string]int `json:"allocation"`
	Sectors    map[string]int `json:"sectors"`
}

type UserInput struct {
	Answers []int `json:"answers"`
}

type UserPortfolio struct {
	UserID        int    `json:"user_id"`
	PortfolioName string `json:"portfolio_name"`
}

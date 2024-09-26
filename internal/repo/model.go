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

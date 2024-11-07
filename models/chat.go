package models

type SubmitChatRequest struct {
	Message []SubmitChatCompletionMessage `json:"messages"`
}

type SubmitChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Holding struct {
	TradingSymbol      string  `json:"tradingsymbol"`
	Exchange           string  `json:"exchange"`
	ISIN               string  `json:"isin"`
	T1Quantity         int     `json:"t1quantity"`
	RealisedQuantity   int     `json:"realisedquantity"`
	Quantity           int     `json:"quantity"`
	AuthorisedQuantity int     `json:"authorisedquantity"`
	Product            string  `json:"product"`
	CollateralQuantity *int    `json:"collateralquantity"`
	CollateralType     *string `json:"collateraltype"`
	Haircut            float64 `json:"haircut"`
	AveragePrice       float64 `json:"averageprice"`
	LTP                float64 `json:"ltp"`
	SymbolToken        string  `json:"symboltoken"`
	Close              float64 `json:"close"`
	ProfitAndLoss      float64 `json:"profitandloss"`
	PNLPercentage      float64 `json:"pnlpercentage"`
}

type TotalHolding struct {
	TotalHoldingValue  float64 `json:"totalholdingvalue"`
	TotalInvValue      float64 `json:"totalinvvalue"`
	TotalProfitAndLoss float64 `json:"totalprofitandloss"`
	TotalPNLPercentage float64 `json:"totalpnlpercentage"`
}

type PortfolioData struct {
	Holdings     []Holding    `json:"holdings"`
	TotalHolding TotalHolding `json:"totalholding"`
}

type PortfolioResponse struct {
	Status    bool          `json:"status"`
	Message   string        `json:"message"`
	ErrorCode string        `json:"errorcode"`
	Data      PortfolioData `json:"data"`
}

type IncomeStatementResponse struct {
	CalendarYear      string `json:"calendarYear"`
	Revenue           int64  `json:"revenue"`
	GrossProfit       int64  `json:"grossProfit"`
	OperatingExpenses int64  `json:"operatingExpenses"`
	OperatingIncome   int64  `json:"operatingIncome"`
	NetIncome         int64  `json:"netIncome"`
}

type BalanceSheetResponse struct {
	CalendarYear     string `json:"calendarYear"`
	TotalAssets      int64  `json:"totalAssets"`
	TotalLiabilities int64  `json:"totalLiabilities"`
}

type FinancialRatiosResponse struct {
	CalendarYear       string  `json:"calendarYear"`
	PriceEarningsRatio float64 `json:"priceEarningsRatio"`
	DebtEquityRatio    float64 `json:"debtEquityRatio"`
	ReturnOnEquity     float64 `json:"returnOnEquity"`
	CurrentRatio       float64 `json:"currentRatio"`
	GrossProfitMargin  float64 `json:"grossProfitMargin"`
}

type CashFlowResponse struct {
	CalendarYear      string `json:"calendarYear"`
	OperatingCashFlow int64  `json:"operatingCashFlow"`
	Inventory         int64  `json:"inventory"`
	FreeCashFlow      int64  `json:"freeCashFlow"`
}

type StockPrice struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	Change float64 `json:"change"`
}

type StockPriceResponse struct {
	Symbol     string       `json:"symbol"`
	Historical []StockPrice `json:"historical"`
}

type FundamentalDataResponse struct {
	BalanceSheetResponse    []*BalanceSheetResponse    `json:"balance_sheet"`
	IncomeStatementResponse []*IncomeStatementResponse `json:"income_statement"`
	CashFlowResponse        []*CashFlowResponse        `json:"Cash_flow"`
	FinancialRatiosResponse []*FinancialRatiosResponse `json:"financial_ratios"`
}

type NewsDataResponse struct {
	Results []NewsResult `json:"results"`
}

type NewsResult struct {
	Insights []NewsDataInsight `json:"insights"`
}

type NewsDataInsight struct {
	Ticker             string `json:"ticker"`
	Sentiment          string `json:"sentiment"`
	SentimentReasoning string `json:"sentiment_reasoning"`
}

type FileUploadRequest struct {
	FilePaths []string `json:"paths"`
}

type FileInfo struct {
	Name string
	Path string
}

type StockResearchReport struct {
	Company            string    `json:"company"`
	Sector             string    `json:"sector"`
	Recommendation     string    `json:"recommendation"`
	TargetPrice        float64   `json:"target_price"`
	RevenueProjections []float64 `json:"revenue_projections"`
	CAGR               float64   `json:"cagr"`
	EBITDA             float64   `json:"ebitda"`
	NewsSummary        string    `json:"news_summary"`
}

type StockResearchResponse struct {
	Data []StockResearchReport `json:"data"`
	Err  string                `json:"err"`
}

type NewsAPIResponse struct {
	Feed []NewsItem `json:"feed"`
}

type NewsItem struct {
	Title           string            `json:"title"`
	TimePublished   string            `json:"time_published"`
	TickerSentiment []TickerSentiment `json:"ticker_sentiment"`
}

type TickerSentiment struct {
	Ticker               string `json:"ticker"`
	RelevanceScore       string `json:"relevance_score"`
	TickerSentimentScore string `json:"ticker_sentiment_score"`
	TickerSentimentLabel string `json:"ticker_sentiment_label"`
}

type FilteredNewsItem struct {
	Title              string `json:"title"`
	TimePublished      string `json:"time_published"`
	Sentiment          string `json:"sentiment"`
	SentimentRelevance string `json:"sentiment_relevance"`
	SentimentLabel     string `json:"sentiment_label"`
}

type UserParseQuery struct {
	Intent      string `json:"intent"`
	Ticker      string `json:"ticker"`
	Amount      string `json:"amount"`
	Horizon     string `json:"horizon"`
	Sector      string `json:"sector"`
	CompanyName string `json:"company_name"`
	News        string `json:"news"`
}

type UserParseQueryResponse struct {
	Data  UserParseQuery `json:"data"`
	Error error          `json:"error"`
}

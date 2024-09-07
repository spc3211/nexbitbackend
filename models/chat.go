package models

type SubmitChatRequest struct {
	Message string `json:"message"`
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
	StockPrice              []*StockPrice              `json:"stock_prices"`
}

package models

type SubmitChatRequest struct {
	Message string `json:"message"`
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
	Status    bool   		  `json:"status"`
	Message   string 		  `json:"message"`
	ErrorCode string 		  `json:"errorcode"`
	Data      PortfolioData   `json:"data"`
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

package util

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

type Data struct {
	Holdings     []Holding    `json:"holdings"`
	TotalHolding TotalHolding `json:"totalholding"`
}

type Response struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorcode"`
	Data      Data   `json:"data"`
}

// Define the constant JSON response
var ConstantResponse = Response{
	Status:    true,
	Message:   "SUCCESS",
	ErrorCode: "",
	Data: Data{
		Holdings: []Holding{
			{
				TradingSymbol:      "AAPL",
				Exchange:           "NASDAQ",
				ISIN:               "US0378331005",
				T1Quantity:         0,
				RealisedQuantity:   10,
				Quantity:           10,
				AuthorisedQuantity: 0,
				Product:            "DELIVERY",
				CollateralQuantity: nil,
				CollateralType:     nil,
				Haircut:            0,
				AveragePrice:       150.00,
				LTP:                175.50,
				SymbolToken:        "AAPL",
				Close:              174.00,
				ProfitAndLoss:      255.00,
				PNLPercentage:      17.00,
			},
			{
				TradingSymbol:      "MSFT",
				Exchange:           "NASDAQ",
				ISIN:               "US5949181045",
				T1Quantity:         0,
				RealisedQuantity:   5,
				Quantity:           5,
				AuthorisedQuantity: 0,
				Product:            "DELIVERY",
				CollateralQuantity: nil,
				CollateralType:     nil,
				Haircut:            0,
				AveragePrice:       290.00,
				LTP:                310.75,
				SymbolToken:        "MSFT",
				Close:              309.50,
				ProfitAndLoss:      103.75,
				PNLPercentage:      7.15,
			},
			{
				TradingSymbol:      "TSLA",
				Exchange:           "NASDAQ",
				ISIN:               "US88160R1014",
				T1Quantity:         0,
				RealisedQuantity:   3,
				Quantity:           3,
				AuthorisedQuantity: 0,
				Product:            "DELIVERY",
				CollateralQuantity: nil,
				CollateralType:     nil,
				Haircut:            0,
				AveragePrice:       750.00,
				LTP:                840.00,
				SymbolToken:        "TSLA",
				Close:              835.00,
				ProfitAndLoss:      270.00,
				PNLPercentage:      12.00,
			},
			{
				TradingSymbol:      "AMZN",
				Exchange:           "NASDAQ",
				ISIN:               "US0231351067",
				T1Quantity:         0,
				RealisedQuantity:   4,
				Quantity:           4,
				AuthorisedQuantity: 0,
				Product:            "DELIVERY",
				CollateralQuantity: nil,
				CollateralType:     nil,
				Haircut:            0,
				AveragePrice:       3400.00,
				LTP:                3550.00,
				SymbolToken:        "AMZN",
				Close:              3540.00,
				ProfitAndLoss:      600.00,
				PNLPercentage:      4.41,
			},
			{
				TradingSymbol:      "NVAX",
				Exchange:           "NASDAQ",
				ISIN:               "US6700024010",
				T1Quantity:         0,
				RealisedQuantity:   12,
				Quantity:           12,
				AuthorisedQuantity: 0,
				Product:            "DELIVERY",
				CollateralQuantity: nil,
				CollateralType:     nil,
				Haircut:            0,
				AveragePrice:       80.00,
				LTP:                90.50,
				SymbolToken:        "NVAX",
				Close:              89.75,
				ProfitAndLoss:      126.00,
				PNLPercentage:      13.12,
			},
			{
				TradingSymbol:      "PLUG",
				Exchange:           "NASDAQ",
				ISIN:               "US72919P2020",
				T1Quantity:         0,
				RealisedQuantity:   25,
				Quantity:           25,
				AuthorisedQuantity: 0,
				Product:            "DELIVERY",
				CollateralQuantity: nil,
				CollateralType:     nil,
				Haircut:            0,
				AveragePrice:       25.00,
				LTP:                32.50,
				SymbolToken:        "PLUG",
				Close:              32.00,
				ProfitAndLoss:      187.50,
				PNLPercentage:      30.00,
			},
			{
				TradingSymbol:      "SOFI",
				Exchange:           "NASDAQ",
				ISIN:               "US83406F1021",
				T1Quantity:         0,
				RealisedQuantity:   30,
				Quantity:           30,
				AuthorisedQuantity: 0,
				Product:            "DELIVERY",
				CollateralQuantity: nil,
				CollateralType:     nil,
				Haircut:            0,
				AveragePrice:       8.00,
				LTP:                9.25,
				SymbolToken:        "SOFI",
				Close:              9.15,
				ProfitAndLoss:      37.50,
				PNLPercentage:      15.63,
			},
		},
		TotalHolding: TotalHolding{
			TotalHoldingValue:  21550.25,
			TotalInvValue:      19900.00,
			TotalProfitAndLoss: 1650.25,
			TotalPNLPercentage: 8.29,
		},
	},
}

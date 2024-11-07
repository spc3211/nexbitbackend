package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBService struct {
	*sqlx.DB
}

func NewDBService(connStr string) (*DBService, error) {
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DBService{db}, nil
}

func (psql *DBService) SaveStockReport(ctx context.Context, report StockResearchReport) error {
	result, err := psql.NamedExec(`
	INSERT 
	INTO 
	stock_reports 
	(company, sector, recommendation, target_price, cagr, ebitda, news_summary) 
	VALUES 
	(:company, :sector, :recommendation, :target_price, :cagr, :ebitda, :news_summary)`, report)
	if err != nil {
		return err // Return the error if the execution fails
	}

	// Check if the row was affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows were inserted")
	}

	return nil

}

func (psql *DBService) FetchStockReport(ctx context.Context, param StockResearchFetchRequest) ([]StockResearchReport, error) {
	query := `
	SELECT 
		id,
    	company,
    	sector,
    	recommendation,
    	target_price,
    	revenue_projections,
		COALESCE(cagr, 0.0) AS cagr,
    	ebitda,
    	ticker,
    	date,
    	news_summary
	FROM 
		stock_reports
	WHERE
	`

	var conditions []string
	var values []interface{}

	var resp []StockResearchReport

	if param.Ticker != "" {
		conditions = append(conditions, "ticker = ?")
		values = append(values, param.Ticker)
	}

	if param.Date != "" {
		parsedDate, err := time.Parse("02/01/2006", param.Date)
		if err != nil {
			return resp, fmt.Errorf("invalid date format: %v", err)
		}

		oneMonthAgo := parsedDate.AddDate(0, -1, 0)
		conditions = append(conditions, "TO_DATE(date, 'DD/MM/YYYY') BETWEEN ? AND ?")
		values = append(values, oneMonthAgo.Format("2006-01-02"), parsedDate.Format("2006-01-02"))
	}

	// Construct the final WHERE clause
	if len(conditions) > 0 {
		query += strings.Join(conditions, " AND ")
	} else {
		return resp, fmt.Errorf("[FetchStockReport] insufficient fetch conditions for fetching reports by params, param: %+v", param)
	}

	query = psql.DB.Rebind(query)

	err := psql.DB.Select(&resp, query, values...)
	if err != nil {
		return resp, fmt.Errorf("[FetchStockReport] failed to fetch stock reports req. err: %v, param: %+v", err, param)
	}
	return resp, err
}

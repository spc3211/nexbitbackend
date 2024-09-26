package repo

import (
	"context"
	"errors"

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

func (dbs *DBService) SaveStockReport(ctx context.Context, report StockResearchReport) error {
	result, err := dbs.NamedExec(`
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

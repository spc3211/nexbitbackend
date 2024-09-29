package service

import (
	"nexbit/internal/repo"
	"nexbit/models"

	"errors"

	"github.com/gofiber/fiber/v2"
)

type OnboardingService struct {
	db       *repo.DBService
}

func NewOnboardingService(db *repo.DBService) *OnboardingService {
	return &OnboardingService{
		db: db,
	}
}

func (s *OnboardingService) SaveUserDetails(ctx *fiber.Ctx, user models.User) error {
	
	result, err := s.db.NamedExec(`
	INSERT
	INTO
	user_details
	(name, email, user_uuid)
	VALUES
	(:name, :email, :user_uuid)`, user)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows were inserted")
	}	

	return nil
}

func (s *OnboardingService) SaveUserPreferences(ctx *fiber.Ctx, preferences []models.UserPreferences) error {

	for _, preference := range preferences {
		result, err := s.db.NamedExec(`
			INSERT
			INTO
			user_preferences
			(user_id, question_id, answer_id)
			VALUES
			(:user_id, :question_id, :answer_id)`, preference)
			if err != nil {
				return err
			}	
		
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return err
			}
			if rowsAffected == 0 {
				return errors.New("no rows were inserted")
			}
		}

	return nil
}

func (s *OnboardingService) SaveUserPortfolio(ctx *fiber.Ctx, userID int, portfolio int, riskScore float64) error {
    result, err := s.db.NamedExec(`
        INSERT INTO user_portfolio
        (user_id, model_portfolio_id, risk_score)
        VALUES (:user_id, :portfolio, :risk_score)`,
        map[string]interface{}{
            "user_id":   userID,       
            "portfolio": portfolio,    
            "risk_score": riskScore,  
        })

    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return errors.New("no rows were inserted")
    }

    return nil
}
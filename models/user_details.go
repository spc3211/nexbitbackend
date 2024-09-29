package models

type User struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
	UserUuid string `json:"user_uuid" db:"user_uuid"`
}

type UserPreferencesInput struct {
    UserID     int `json:"user_id"`
    Preferences []struct {
        QuestionID int `json:"question_id"`
        AnswerID   int `json:"answer_id"`
    } `json:"preferences"`
}

type UserPreferences struct {
    UserID     int `json:"user_id" db:"user_id"`
    QuestionID int `json:"question_id" db:"question_id"`
    AnswerID   int `json:"answer_id" db:"answer_id"`
}

type UserPortfolio struct {
	ID               int `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID           int `json:"user_id" db:"user_id"`
	ModelPortfolioID int `json:"model_portfolio_id" db:"model_portfolio_id"`
	RiskScore        int `json:"risk_score" db:"risk_score"`
}
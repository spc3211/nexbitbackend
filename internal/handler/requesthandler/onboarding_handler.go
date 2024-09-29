package requesthandler

import (
	"math"
	models "nexbit/models"
	"nexbit/internal/constants"
	"github.com/gofiber/fiber/v2"

	onboardingService "nexbit/internal/service"
	"github.com/google/uuid"
)

type OnboardingHandler struct {
	onboardingService *onboardingService.OnboardingService
}

func NewOnboardingHandler(onboardingService *onboardingService.OnboardingService) *OnboardingHandler {
	return &OnboardingHandler{
		onboardingService: onboardingService,
	}
}

func (h *OnboardingHandler) SaveUserDetails(c *fiber.Ctx) error {
	var user models.User
	//generate and add UserUuid
	userUUID := uuid.New()
	user.UserUuid = userUUID.String()
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.onboardingService.SaveUserDetails(c, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User details saved successfully"})
}

func (h *OnboardingHandler) SaveUserPreferences(c *fiber.Ctx) error {
    var input models.UserPreferencesInput
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Convert input to a list of UserPreferences for saving
    var preferences []models.UserPreferences
	var userAnswers []int
    for _, pref := range input.Preferences {
        preferences = append(preferences, models.UserPreferences{
            UserID:     input.UserID,
            QuestionID: pref.QuestionID,
            AnswerID:   pref.AnswerID,
        })

		answerScore := getAnswerScore(pref.QuestionID, pref.AnswerID)
        userAnswers = append(userAnswers, answerScore)
    }

    if err := h.onboardingService.SaveUserPreferences(c, preferences); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

	riskScore := calculateRiskScore(userAnswers)
	allocatedPortfolio, normalizedScore := allocatePortfolio(riskScore)

	if err := h.onboardingService.SaveUserPortfolio(c, input.UserID, allocatedPortfolio, normalizedScore); err != nil {
    	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}


    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "User preferences saved and portfolio allocated successfully",
        "portfolio": allocatedPortfolio,
    })
}

func calculateRiskScore(userAnswers []int) int {
    riskScore := 0
    for _, score := range userAnswers {
        riskScore += score
    }
    return riskScore
}

func allocatePortfolio(riskScore int) (int, float64) {
    portfolios := constants.Portfolios
    normalizedScore := (float64(riskScore - 4) / 8.0)
    portfolioIndex := int(math.Min(math.Max(float64(normalizedScore*float64(len(portfolios))), 0), float64(len(portfolios)-1)))
    return portfolios[portfolioIndex].ID, normalizedScore
}

func getAnswerScore(questionID int, answerID int) int {
    for _, question := range constants.Questions {
        if question.ID == questionID {
            for _, answer := range question.Answers {
                if answer.Id == answerID {
                    return answer.Score
                }
            }
        }
    }
    return 0
}
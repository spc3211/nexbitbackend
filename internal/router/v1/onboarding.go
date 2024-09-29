package v1

import (
	requesthandler "nexbit/internal/handler/requesthandler"
	onboarding "nexbit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func OnboardingRouter(app *fiber.App, onboardingService *onboarding.OnboardingService) {
	onboardingHandler := requesthandler.NewOnboardingHandler(onboardingService)
	api := app.Group("/v1")
	api.Post("/user", onboardingHandler.SaveUserDetails)
	api.Post("/user/preferences", onboardingHandler.SaveUserPreferences)
}
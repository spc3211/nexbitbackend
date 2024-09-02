package handler

import (
	chatService "nexbit/internal/service"
	"nexbit/util"

	models "nexbit/models"

	"github.com/gofiber/fiber/v2"
)

type ChatHandler struct {
	authService *chatService.ChatService
}

func NewChatHandler(authService *chatService.ChatService) *ChatHandler {
	return &ChatHandler{authService: authService}
}

func (h *ChatHandler) ChatHandler(ctx *fiber.Ctx) error {

	var reqData models.SubmitChatRequest

	err := ctx.BodyParser(&reqData)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[ChatHandler] request body is invalid with err:%v", err)
		return err
	}

	ctx.Locals("requestData", reqData)
	err = h.authService.ChatService(ctx)
	return ctx.Context().Err()
}

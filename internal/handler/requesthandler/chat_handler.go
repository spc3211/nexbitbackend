package requesthandler

import (
	chatService "nexbit/internal/service"
	"nexbit/util"

	models "nexbit/models"

	"github.com/gofiber/fiber/v2"
)

type ChatHandler struct {
	chatService *chatService.ChatService
}

func NewChatHandler(chatService *chatService.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

func (h *ChatHandler) UserChatHandler(ctx *fiber.Ctx) error {

	var reqData models.SubmitChatRequest

	err := ctx.BodyParser(&reqData)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[ChatHandler] request body is invalid with err:%v", err)
		return err
	}

	ctx.Locals("requestData", reqData)
	err = h.chatService.ChatService(ctx)
	return ctx.Context().Err()
}

func (h *ChatHandler) FundamentalHandler(ctx *fiber.Ctx) error {

	stockSymbol := ctx.Query("stock")
	ctx.Locals("stockSymbol", stockSymbol)
	_ = h.chatService.FetchFundamentals(ctx)
	return ctx.Context().Err()
}

func (h *ChatHandler) NewsInsightsHandler(ctx *fiber.Ctx) error {
	stockSymbol := ctx.Query("stock")
	ctx.Locals("stockSymbol", stockSymbol)
	_ = h.chatService.FetchNewsInsights(ctx)
	return ctx.Context().Err()
}
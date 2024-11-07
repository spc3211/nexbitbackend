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
	return &ChatHandler{
		chatService: chatService,
	}
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
	return err
}

func (h *ChatHandler) FundamentalHandler(ctx *fiber.Ctx) error {

	stockSymbol := ctx.Query("stock")
	ctx.Locals("stockSymbol", stockSymbol)
	finalRespnse, _ := h.chatService.FetchFundamentals(ctx, stockSymbol)
	return ctx.JSON(fiber.Map{
		"stock_financials": finalRespnse,
	})
}

func (h *ChatHandler) NewsInsightsHandler(ctx *fiber.Ctx) error {
	stockSymbol := ctx.Query("stock")
	ctx.Locals("stockSymbol", stockSymbol)
	_ = h.chatService.FetchNewsInsights(ctx)
	return ctx.Context().Err()
}

func (h *ChatHandler) FileUploadHandler(ctx *fiber.Ctx) error {
	var reqData models.FileUploadRequest

	err := ctx.BodyParser(&reqData)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[FileUploadHandler] request body is invalid with err:%v", err)
		return err
	}

	err = h.chatService.Uploadfile(ctx, reqData)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[FileUploadHandler] error from fileUploadService err:%v", err)
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "uploaded succesfully",
	})
}

func (h *ChatHandler) UserQueryHandler(ctx *fiber.Ctx) error {

	var reqData models.SubmitChatRequest

	err := ctx.BodyParser(&reqData)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[UserQueryHandler] request body is invalid with err:%v", err)
		return err
	}

	chatResponse, err := h.chatService.UserQueryService(ctx, reqData)
	if err != nil {
		util.WithContext(ctx.Context()).Errorf("[UserQueryHandler] Failed to process chat request. err: %v", err)
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": chatResponse,
	})
}

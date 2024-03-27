package controllers

import (
	"jasanfarah/chat-service/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ConversationControllerInterface interface {
	CreateConversation(ctx *fiber.Ctx) error
	
}

type ConversationControllerOptions struct {
	ConversationService services.ConversationServiceInterface
	Logger              *zap.Logger
	Validator           *validator.Validate
}

type ConversationController struct {
	service   services.ConversationServiceInterface
	logger    *zap.Logger
	validator *validator.Validate
}
func NewConversationController(options ConversationControllerOptions) ConversationControllerInterface {
	return &ConversationController{
		service:   options.ConversationService,
		logger:    options.Logger,
		validator: options.Validator,
	}
}

// CreateConversation implements ConversationControllerInterface.
func (c *ConversationController) CreateConversation(ctx *fiber.Ctx) error {

	payload := services.CreateConversationInput{}
	if err := ctx.BodyParser(&payload); err != nil {
		c.logger.Error("Failed to parse request body", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
		})
	}

	err := c.validator.Struct(payload)
	if err != nil {
		c.logger.Error("Failed to validate request body", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to validate request body",
		})

	}

	output, err := c.service.CreateConversation(payload)
if err != nil {
return err
}

return ctx.Status(fiber.StatusCreated).JSON(output)
}




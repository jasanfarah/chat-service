package controllers

import (
	"fmt"
	createmessageDTO "jasanfarah/chat-service/internal/api/dto/createmessage"
	"jasanfarah/chat-service/internal/models"
	"jasanfarah/chat-service/internal/services"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"go.uber.org/zap"
)

type MessageControllerInterface interface {
	CreateMessage(ctx *fiber.Ctx) error
	GetMessages(ctx *fiber.Ctx) error
}

type MessageControllerOptions struct {
	MessageService      services.MessageServiceInterface
	ConversationService services.ConversationServiceInterface
	Logger              *zap.Logger
	Validation          *validator.Validate
}

type MessageController struct {
	service             services.MessageServiceInterface
	conversationService services.ConversationServiceInterface
	logger              *zap.Logger
	validation          *validator.Validate
}

func NewMessageController(options MessageControllerOptions) MessageControllerInterface {
	return &MessageController{
		service:             options.MessageService,
		conversationService: options.ConversationService,
		logger:              options.Logger,
		validation:          options.Validation,
	}

}

// CreateMessage implements MessageControllerInterface.

/*
Flow:

1. Parse the request body
2. Validate the request body
3. Create a new message
4. Call AI agent
5. Update the conversation's last message at


*/
func (c *MessageController) CreateMessage(ctx *fiber.Ctx) error {

	payload := createmessageDTO.CreateMessageInput{}
	if err := ctx.BodyParser(&payload); err != nil {
		c.logger.Error("Failed to parse request body", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
		})
	}

	err := c.validation.Struct(payload)
	if err != nil {
		c.logger.Error("Failed to validate request body", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to validate request body",
		})
	}

	message := &models.Message{
		ID:             uuid.New(),
		ConversationID: payload.ConversationID,
		Content:        payload.Messages[len(payload.Messages)-1].Content,
		Role:           payload.Messages[len(payload.Messages)-1].Role,
		Embedding:      pgvector.NewVector([]float32{0.1, 0.2, 0.3, 0.4, 0.5}),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	fmt.Println(message)

	message, err = c.service.CreateMessage(message)
	if err != nil {
		c.logger.Error("Failed to create message", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create message",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(message)
}


// GetMessages implements MessageControllerInterface.
func (m *MessageController) GetMessages(ctx *fiber.Ctx) error {
    payload :=services.GetMessagesInput{}
    if err := ctx.BodyParser(&payload); err != nil {
		m.logger.Error("Failed to parse request body", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
		})
	}


    messages, err := m.service.GetMessages(&payload)
    if err != nil {
        m.logger.Error("Failed to get messages", zap.Error(err))
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to get messages",
        })
    }


    return ctx.Status(fiber.StatusOK).JSON(messages)


}

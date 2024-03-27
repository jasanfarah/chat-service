package api

import (
	"jasanfarah/chat-service/internal/api/controllers"
	"jasanfarah/chat-service/internal/api/groups"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func InitializeAPIRoutes(app *fiber.App, logger *zap.Logger,conversationController controllers.ConversationControllerInterface , messageController controllers.MessageControllerInterface, validation *validator.Validate) {

    app.Use(recover.New())
	app.Use(healthcheck.New())

	apiGroup := app.Group("/v1/api")

  
    groups.InitializeConversationRoutes(apiGroup, conversationController, messageController, logger)


}


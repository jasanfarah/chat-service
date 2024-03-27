package groups

import (
	"jasanfarah/chat-service/internal/api/controllers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)


func InitializeConversationRoutes(app fiber.Router, controller controllers.ConversationControllerInterface, messageController controllers.MessageControllerInterface, logger *zap.Logger) {
    conversationGroup := app.Group("/conversation")
    

  conversationGroup.Post("/", controller.CreateConversation)
  conversationGroup.Post("/:id/message", messageController.CreateMessage)
  conversationGroup.Get("/:id/message", messageController.GetMessages)

    

    
    
    
}
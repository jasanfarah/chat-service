package services

import (
	"jasanfarah/chat-service/internal/models"
	"jasanfarah/chat-service/internal/repositories"
	"time"

	"go.uber.org/zap"
)

type MessageServiceInterface interface {
	CreateMessage(input *models.Message) (*models.Message, error)
	GetMessages(input *GetMessagesInput) ([]GetMessagesResponse, error)
}

type MessageServiceOptions struct {
	Repository            repositories.MessageRepositoryInterface
	Logger                *zap.Logger
	ConversationRepostory repositories.ConversationRepositoryInterface
}

type MessageService struct {
	repository            repositories.MessageRepositoryInterface
	logger                *zap.Logger
	conversationRepostory repositories.ConversationRepositoryInterface
}

func NewMessageService(options MessageServiceOptions) MessageServiceInterface {
	return &MessageService{
		repository:            options.Repository,
		logger:                options.Logger,
		conversationRepostory: options.ConversationRepostory,
	}
}

type CreateMessageOutput struct {
	Message *models.Message `json:"message"`
}

// CreateMessage implements MessageServiceInterface.
func (m *MessageService) CreateMessage(input *models.Message) (*models.Message, error) {
	conversation, err := m.conversationRepostory.GetConversationByID(input.ConversationID.String())
	if err != nil {
		return nil, err
	}
	conversation.UpdatedAt = time.Now()

	err = m.repository.CreateMessage(input)
	if err != nil {
		m.logger.Error("Failed to create message", zap.Error(err))
		return nil, err
	}

	err = m.conversationRepostory.UpdateConversation(conversation)
	if err != nil {
		m.logger.Error("Failed to update conversation", zap.Error(err))
		return nil, err
	}

	return input, nil

}

type GetMessagesInput struct {
	ConversationID string `json:"conversation_id"`
}

type Message struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type GetMessagesResponse struct {
	Messages []Message `json:"messages"`
}

func (m *MessageService) GetMessages(input *GetMessagesInput) ([]GetMessagesResponse, error) {

	_, err := m.conversationRepostory.GetConversationByID(input.ConversationID)
	if err != nil {
		return nil, err
	}
	messages, err := m.repository.GetMessages(input.ConversationID)
	if err != nil {
		m.logger.Error("Failed to get messages", zap.Error(err))
		return nil, err
	}

	var filteredMessages []Message
	for _, msg := range messages {
		if msg.ConversationID.String() == input.ConversationID {
			filteredMessages = append(filteredMessages, Message{
				ID:        msg.ID.String(),
				Content:   msg.Content,
				Role:      msg.Role,
				CreatedAt: msg.CreatedAt,
			})
		}
	}

	response := GetMessagesResponse{
		Messages: filteredMessages,
	}

	return []GetMessagesResponse{response}, nil

}

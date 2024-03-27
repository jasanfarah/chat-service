package services

import (
	"jasanfarah/chat-service/internal/models"
	"jasanfarah/chat-service/internal/repositories"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ConversationServiceInterface interface {
	CreateConversation(createConversationInput CreateConversationInput) (CreateConversationOutput, error)
	GetConversationByID(getConversationByIDInput GetConversationByIDInput) (GetConversationByIDOutput, error)
	GetConversations() (GetConversationsOutput, error)
	UpdateConversation(updateConversationInput UpdateConversationInput) (UpdateConversationOutput, error)
}

type ConversationService struct {
	repository repositories.ConversationRepositoryInterface
	log        *zap.Logger
}


type ConversationServiceOptions struct {
	Repository repositories.ConversationRepositoryInterface
	Logger     *zap.Logger
}

func NewConversationService(options ConversationServiceOptions) ConversationServiceInterface {
	return &ConversationService{
		repository: options.Repository,
		log:        options.Logger,
	}
}





// CreateConversation implements ConversationServiceInterface.

type CreateConversationInput struct {
}

type CreateConversationOutput struct {
	Conversation *models.Conversation `json:"conversation"`
}


func (c *ConversationService) CreateConversation(createConversationInput CreateConversationInput) (CreateConversationOutput, error) {
        
        conversation := &models.Conversation{
            ID: uuid.New(),
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        }
        
        err := c.repository.CreateConversation(conversation)
        if err != nil {
            c.log.Error("Failed to create conversation", zap.Error(err))
            return CreateConversationOutput{}, err
        }
        
        return CreateConversationOutput{
            Conversation: conversation,
        }, nil

        
}


type GetConversationByIDInput struct {
	ID string `json:"id" validate:"required"`
}

type GetConversationByIDOutput struct {
	Conversation *models.Conversation `json:"conversation"`
}


// GetConversationByID implements ConversationServiceInterface.
func (c *ConversationService) GetConversationByID(getConversationByIDInput GetConversationByIDInput) (GetConversationByIDOutput, error) {
	panic("unimplemented")
}


type GetConversationsInput struct {
}

type GetConversationsOutput struct {
	Conversations []models.Conversation `json:"conversations"`
}


// GetConversations implements ConversationServiceInterface.
func (c *ConversationService) GetConversations() (GetConversationsOutput, error) {
	panic("unimplemented")
}

type UpdateConversationInput struct {
	ID string `json:"id" validate:"required"`
}

type UpdateConversationOutput struct {
	Conversation *models.Conversation `json:"conversation"`
}

func (c *ConversationService) UpdateConversation(updateConversationInput UpdateConversationInput) (UpdateConversationOutput, error) {
	conversation, err := c.repository.GetConversationByID(updateConversationInput.ID)
	if err != nil {
		c.log.Error("Failed to get conversation by id", zap.Error(err))
		return UpdateConversationOutput{}, err
	}

	conversation.UpdatedAt = time.Now()

	err = c.repository.UpdateConversation(conversation)
	if err != nil {
		c.log.Error("Failed to update conversation", zap.Error(err))
		return UpdateConversationOutput{}, err
	}

	return UpdateConversationOutput{
		Conversation: conversation,
	}, nil
	
}


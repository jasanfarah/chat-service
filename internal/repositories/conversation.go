package repositories

import (
    "jasanfarah/chat-service/internal/models"
    "gorm.io/gorm"
    "go.uber.org/zap"
)

type ConversationRepositoryInterface interface {
    CreateConversation(conversation *models.Conversation) error
    UpdateConversation(conversation *models.Conversation) error
    GetConversationByID(id string) (*models.Conversation, error)
    GetConversations() ([]models.Conversation, error)

}


type ConversationRepository struct {
    db *gorm.DB
    logger *zap.Logger
}



func NewConversationRepository(db *gorm.DB, logger *zap.Logger) *ConversationRepository {
    return &ConversationRepository{
        db: db,
        logger: logger,
    }
}

func (r *ConversationRepository) CreateConversation(conversation *models.Conversation) error {
    err := r.db.Create(conversation).Error
    if err != nil {
        r.logger.Error("Failed to create conversation", zap.Error(err))
        return err
    }
    return nil
}

func (r *ConversationRepository) GetConversationByID(id string) (*models.Conversation, error) {

    conversation := &models.Conversation{}
    err := r.db.Where("id = ?", id).First(conversation).Error
    if err != nil {
        r.logger.Error("Failed to get conversation by id", zap.Error(err))
        return nil, err
    }
    return conversation, nil
}


func (r *ConversationRepository) GetConversations() ([]models.Conversation, error) {
    var conversations []models.Conversation
    err := r.db.Find(&conversations).Error
    if err != nil {
        r.logger.Error("Failed to get conversations", zap.Error(err))
        return nil, err
    }
    return conversations, nil
}


func (r *ConversationRepository) UpdateConversation(conversation *models.Conversation) error {
    err := r.db.Save(conversation).Error
    if err != nil {
        r.logger.Error("Failed to update conversation", zap.Error(err))
        return err
    }
    return nil
}


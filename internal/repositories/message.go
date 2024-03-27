package repositories

import (
    "jasanfarah/chat-service/internal/models"
    "gorm.io/gorm"
    "go.uber.org/zap"
)

type MessageRepositoryInterface interface {
    CreateMessage(message *models.Message) error
    GetMessages(conversationID string) ([]models.Message, error)
}

type MessageRepository struct {
    db *gorm.DB
    logger *zap.Logger
}

func NewMessageRepository(db *gorm.DB, logger *zap.Logger) *MessageRepository {
    return &MessageRepository{
        db: db,
        logger: logger,
    }
}


func (r *MessageRepository) CreateMessage(message *models.Message) error {
    err := r.db.Create(message).Error
    if err != nil {
        r.logger.Error("Failed to create message", zap.Error(err))
        return err
    }
    return nil
}


func (r *MessageRepository) GetMessages(conversationID string) ([]models.Message, error) {
    var messages []models.Message
    err := r.db.Where("conversation_id = ?", conversationID).Find(&messages).Error
    if err != nil {
        r.logger.Error("Failed to get messages", zap.Error(err))
        return nil, err
    }
    return messages, nil
}


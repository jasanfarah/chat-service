package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"

)


type Message struct {
	ID             uuid.UUID    `gorm:"primaryKey;type:uuid;"`
	Conversation   Conversation `gorm:"foreignKey:ConversationID;not null"`
	ConversationID uuid.UUID    // Foreign key to associate with the Conversation
	Content        string
	Role           string
	Embedding      pgvector.Vector `pg:"type:vector(768)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}


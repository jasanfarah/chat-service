package models  
import (
    "time"
    "github.com/google/uuid"
)

type Conversation struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


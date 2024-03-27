package createmessageDTO

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

const (
	RoleUser = "user"
	RoleBot  = "agent"
)


type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CreateMessageInput struct {
    ConversationID uuid.UUID `json:"conversation_id" validate:"required"`
    Messages     []Message `json:"messages" validate:"required"`
}


func (m *CreateMessageInput) UnmarshalJSON(data []byte) error {
    type Alias CreateMessageInput
    aux := &struct {
        *Alias
    }{
        Alias: (*Alias)(m),
    }

    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    // Ensure at least one message is from a user
    userMessageFound := false
    for _, message := range m.Messages {
        
        if message.Role == RoleUser {
            userMessageFound = true
            break
        }
    }

    if !userMessageFound {
        return errors.New("at least one message must be from a user")
    }

    return nil
}


type CreateMessageOutput struct {
    Message string `json:"message"`
}



package migrations

import (
    "jasanfarah/chat-service/internal/models"
    "gorm.io/gorm"
)


func InitialMigrate(db *gorm.DB) error {
     
   err:= db.AutoMigrate(&models.Conversation{}, &models.Message{}) 

    if err != nil {
        return err
    }
    return nil

}


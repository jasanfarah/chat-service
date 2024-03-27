package migrations

import (
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
        
    err:= InitialMigrate(db)

        if err != nil {
            return err
        }
        return nil
    
    }



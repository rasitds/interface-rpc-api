package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Theme struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name       string
	Background string
	Foreground string
}

func AutoMigrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := db.AutoMigrate(&Theme{}); err != nil {
		return err
	}
	return nil
}

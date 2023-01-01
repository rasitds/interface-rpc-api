package models

import "gorm.io/gorm"

type Theme struct {
	Background string `json:"background"`
	Foreground string `json:"foreground"`
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&Theme{}); err != nil {
		return err
	}
	return nil
}

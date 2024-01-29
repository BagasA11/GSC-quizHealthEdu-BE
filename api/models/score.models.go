package models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
}

func (score *Score) BeforeCreate(tx *gorm.DB) error {

	return nil
}

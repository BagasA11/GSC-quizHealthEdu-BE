package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
}

// Before
func (quiz *Quiz) BeforeCreate(tx *gorm.DB) error {

	return nil
}

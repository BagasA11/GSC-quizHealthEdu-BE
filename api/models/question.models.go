package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
}

// Before
func (question *Question) BeforeCreate(tx *gorm.DB) error {

	return nil
}

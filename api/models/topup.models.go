package models

import (
	"gorm.io/gorm"
)

type TopUp struct {
	gorm.Model
}

func (topUp *TopUp) BeforeCreate(tx *gorm.DB) error {

	return nil
}

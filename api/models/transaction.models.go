package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
}

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) error {

	return nil
}

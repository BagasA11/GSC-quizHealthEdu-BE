package models

import (
	"gorm.io/gorm"
)

type TopUp struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	Status bool `gorm:"type:boolean; not null; default:true"`
	UserID uint `gorm:"not null;"`
	User   User
	QuizID uint `gorm:"not null;"`
	Quiz   Quiz
}

func (topUp *TopUp) BeforeCreate(tx *gorm.DB) error {

	return nil
}

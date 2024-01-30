package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Question string `gorm:"not null"`
	Answer   string `gorm:"not null;size:1"`
	QuizID   uint
	Quiz     Quiz
	Option   []Option
}

// Before
func (question *Question) BeforeCreate(tx *gorm.DB) error {

	return nil
}

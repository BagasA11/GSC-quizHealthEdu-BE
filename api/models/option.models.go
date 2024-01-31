package models

import (
	"gorm.io/gorm"
)

type Option struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	Alphabet   string `gorm:"not null;size:1"`
	Text       string `gorm:"not null;size:200"`
	Color      string `gorm:"not null;size:30;default:white"`
	QuestionID uint
	Question   Question `gorm:"constraint:OnDelete:CASCADE;"`
}

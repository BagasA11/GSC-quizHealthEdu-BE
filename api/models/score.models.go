package models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	Point     float64 `gorm:"not null;default:0"`
	CreatedAt uint    `gorm:"type:integer; not null"`
	UpdatedAt uint    `gorm:"type:integer; not null"`
	DeletedAt uint    `gorm:"type:integer; default:null"`
	UserID    uint    `gorm:"not null;"`
	User      User
	QuizID    uint `gorm:"not null;"`
	Quiz      Quiz
}

func (score *Score) BeforeCreate(tx *gorm.DB) error {

	return nil
}

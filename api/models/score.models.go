package models

import (
	"time"

	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	Point     float64 `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID    uint `gorm:"not null;"`
	User      User
	QuizID    uint `gorm:"not null;"`
	Quiz      Quiz
}

func (score *Score) BeforeCreate(tx *gorm.DB) error {

	return nil
}

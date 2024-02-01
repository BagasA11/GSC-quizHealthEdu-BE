package models

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null;unique"`
	Topic     string `gorm:"not null"`
	Img       *string
	Free      bool `gorm:"type:boolean; not null; default:true"`
	Price     *uint64
	Disc      uint8 `gorm:"type:integer; not null; default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Score     []Score
	TopUp     []TopUp
	Question  []Question
}

// Before
func (quiz *Quiz) BeforeCreate(tx *gorm.DB) error {

	// generate time
	tx.Statement.SetColumn("CreatedAt", time.Now())
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}

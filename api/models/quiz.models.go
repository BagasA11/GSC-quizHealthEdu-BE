package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null;unique"`
	Topic     string `gorm:"not null"`
	Free      bool   `gorm:"type:boolean; not null; default:true"`
	Price     *uint64
	Disc      uint8 `gorm:"type:integer; not null; default:0"`
	CreatedAt uint  `gorm:"type:integer; not null"`
	UpdatedAt uint  `gorm:"type:integer; not null"`
	DeletedAt uint  `gorm:"type:integer; default:null"`
	Score     []Score
	TopUp     []TopUp
	Question  []Question
}

// Before
func (quiz *Quiz) BeforeCreate(tx *gorm.DB) error {

	return nil
}

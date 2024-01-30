package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Qty       uint64 `gorm:"not null"`
	Pay       bool   `gorm:"type:boolean; not null; default:false"`
	Cancel    bool   `gorm:"type:boolean; not null; default:false"`
	PayAt     uint   `gorm:"type:integer; null"`
	CancelAt  uint   `gorm:"type:integer; null"`
	CreatedAt uint   `gorm:"type:integer; not null"`
	UpdatedAt uint   `gorm:"type:integer; not null"`
	DeletedAt uint   `gorm:"type:integer; default:null"`
	UserID    uint
	User      User
}

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) error {

	return nil
}

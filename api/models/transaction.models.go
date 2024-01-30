package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Qty       uint64 `gorm:"not null"`
	Pay       bool   `gorm:"type:boolean; not null; default:false"`
	Cancel    bool   `gorm:"type:boolean; not null; default:false"`
	PayAt     *time.Time
	CancelAt  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID    uint
	User      User
}

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) error {

	return nil
}

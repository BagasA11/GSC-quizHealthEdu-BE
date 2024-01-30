package models

import (
	"gorm.io/gorm"
)

// "fmt"
// "time"

// "golang.org/x/crypto/bcrypt"
// "gorm.io/gorm"
type User struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"not null;unique;size:100"`
	Email       string `gorm:"not null;unique;size:100"`
	Wallet      uint64 `gorm:"type:integer; not null; default:0"`
	Password    string `gorm:"not null"`
	Bio         *string
	Avatar      *string
	Admin       bool `gorm:"type:boolean; not null; default:false"`
	CreatedAt   uint `gorm:"type:integer; not null"`
	UpdatedAt   uint `gorm:"type:integer; not null"`
	DeletedAt   uint `gorm:"type:integer; default:null"`
	Score       []Score
	TopUp       []TopUp
	Transaction []Transaction
}

// Before
func (User *User) BeforeCreate(tx *gorm.DB) error {

	return nil
}

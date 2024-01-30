package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// "fmt"
// "time"
// golang.org/x/crypto/bcrypt
// "golang.org/x/crypto/bcrypt"
// "gorm.io/gorm"
type User struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"not null;unique;size:100"`
	Email       string `gorm:"not null;unique;size:100"`
	Block       bool   `gorm:"type:boolean; not null; default:false"`
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

// perform action Before Create record
// password must be hashed before store
// created_at and updated_at must be set
// return error if occur
func (user *User) BeforeCreate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	fmt.Print("hash:", hash)
	tx.Statement.SetColumn("Password", hash)

	// generate time
	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

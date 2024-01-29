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
}

// Before
func (User *User) BeforeCreate(tx *gorm.DB) error {

	return nil
}

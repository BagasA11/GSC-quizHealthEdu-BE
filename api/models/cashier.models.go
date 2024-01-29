package models

import (
	"gorm.io/gorm"
)

// "fmt"
// "time"

// "golang.org/x/crypto/bcrypt"
// "gorm.io/gorm"
type Cashier struct {
	gorm.Model
}

// Before
func (cashier *Cashier) BeforeCreate(tx *gorm.DB) error {

	return nil
}

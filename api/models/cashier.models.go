package models

// "fmt"
// "time"
// // "gorm.io/driver/mysql"
// "golang.org/x/crypto/bcrypt"
import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Cashier struct {
	gorm.Model
	ID        uint   `gorm:"type:integer; primaryKey"`
	Username  string `gorm:"type:string; not null; unique; size:20"`
	Email     string `gorm:"type:string; size:30; not null; unique"`
	Contact   string `gorm:"type:string; size:20; null;"`
	Password  string `gorm:"type:string; not null"`
	RetailID  uint
	Retail    Retail
	Sales     []Sale
	CreatedAt uint `gorm:"type:integer; not null"`
	UpdatedAt uint `gorm:"type:integer; not null"`
	DeletedAt uint `gorm:"type:integer; default:null"`
}

// Before
func (cashier *Cashier) BeforeCreate(tx *gorm.DB) error {
	// hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(cashier.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	fmt.Println(cashier.Password)
	tx.Statement.SetColumn("Password", hash)

	// generate time
	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

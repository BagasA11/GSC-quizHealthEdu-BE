package repository

import (
	"BagasA11/GSC-quizHealthEdu-BE/configs"

	"gorm.io/gorm"
)

// "time"
// "BagasA11/GSC-quizHealthEdu-BE/api/models"
type CashierRepository struct {
	Db *gorm.DB
}

/*create Cashier Repository instance*/
func NewCashierRepository() *CashierRepository {
	return &CashierRepository{
		Db: configs.GetDB(),
	}
}

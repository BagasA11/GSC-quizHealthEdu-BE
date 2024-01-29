package repository

import (
	"BagasA11/GSC-quizHealthEdu-BE/configs"

	"gorm.io/gorm"
)

// "time"
// "BagasA11/GSC-quizHealthEdu-BE/api/models"
type UserRepository struct {
	Db *gorm.DB
}

/*create Cashier Repository instance*/
func NewUserRepository() *UserRepository {
	return &UserRepository{
		Db: configs.GetDB(),
	}
}

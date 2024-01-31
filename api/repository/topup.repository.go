package repository

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/models"
	"BagasA11/GSC-quizHealthEdu-BE/configs"

	"gorm.io/gorm"
)

type TopUpRepository struct {
	Db *gorm.DB
}

func NewTopupRepo() *TopUpRepository {
	return &TopUpRepository{
		Db: configs.GetDB(),
	}
}

func (tr *TopUpRepository) Create(topup models.TopUp) error {
	tx := tr.Db.Begin()
	err := tx.Create(&topup).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (tr *TopUpRepository) All() ([]models.TopUp, error) {
	var topups []models.TopUp
	err := tr.Db.Find(&topups).Error
	return topups, err
}

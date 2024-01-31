package repository

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/models"
	"BagasA11/GSC-quizHealthEdu-BE/configs"

	"gorm.io/gorm"
)

type TransRepository struct {
	Db *gorm.DB
}

func NewTransRepository() *TransRepository {
	return &TransRepository{
		Db: configs.GetDB(),
	}
}

func (transRpo *TransRepository) Create(trans models.Transaction) error {
	tx := transRpo.Db.Begin()
	err := tx.Create(&trans).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (transRpo *TransRepository) FindID(id uint) (models.Transaction, error) {
	var trans models.Transaction
	//SELECT * FROM transactions WHERE id = {id}
	err := transRpo.Db.Where("id = ?", id).First(&trans).Error
	return trans, err
}

package repository

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/models"
	"BagasA11/GSC-quizHealthEdu-BE/configs"
	"errors"

	"gorm.io/gorm"
)

type BlacklistRepository struct {
	DB *gorm.DB
}

func NewBlacklistRepository() *BlacklistRepository {
	return &BlacklistRepository{
		DB: configs.GetDB(),
	}
}

func (br *BlacklistRepository) Create(bt models.BlacklistToken) error {
	tx := br.DB.Begin()
	err := tx.Create(&br).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (br *BlacklistRepository) FindToken(oldtoken string) error {
	err := br.DB.Where("token = ?", oldtoken).First(models.BlacklistToken{}).Error
	//if error occurs, that's mean token is have'nt blacklisted yet
	//so token still be valid
	if err != nil {
		return nil
	}
	return errors.New("blacklist token found")
}

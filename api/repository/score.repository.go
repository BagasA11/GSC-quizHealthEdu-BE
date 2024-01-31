package repository

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/models"
	"BagasA11/GSC-quizHealthEdu-BE/configs"
	"time"

	"gorm.io/gorm"
)

type ScoreRepository struct {
	Db *gorm.DB
}

func NewScoreRepository() *ScoreRepository {
	return &ScoreRepository{
		Db: configs.GetDB(),
	}
}

func (sr *ScoreRepository) CreateScore(score models.Score) error {
	tx := sr.Db.Begin()
	err := tx.Create(&score).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (sr *ScoreRepository) Rank() ([]models.Score, error) {
	var ranks []models.Score
	// SELECT * FROM scores ORDER BY point DESC
	err := sr.Db.Find(&ranks).Order("point DESC").Error
	return ranks, err
}

func (sr *ScoreRepository) FindByUserID(id uint) (models.Score, error) {
	var userScore models.Score
	//SELECT * FROM scores WHERE user_id = {id}
	err := sr.Db.Where("user_id = ?", id).First(&userScore).Error
	return userScore, err
}

func (sr *ScoreRepository) Recents() ([]models.Score, error) {
	var scores []models.Score
	//SELECT * FROM scores WHERE created_at >= DATEADD(hour, -2, GETDATE())
	err := sr.Db.Where("created_at >= ?", time.Now().Add(time.Hour*(-2))).Find(&scores).Order("point DESC").Error
	return scores, err
}

func (sr *ScoreRepository) Update(score models.Score) error {
	tx := sr.Db.Begin()
	err := tx.Model(&models.Score{}).Where("id = ?", score.ID).Updates(&score).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (sr *ScoreRepository) Delete(id uint) error {
	tx := sr.Db.Begin()
	err := tx.Model(&models.Score{}).Where("id = ?", id).Update("deleted_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

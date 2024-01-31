package repository

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/models"
	"BagasA11/GSC-quizHealthEdu-BE/configs"
	"time"

	"gorm.io/gorm"
)

type QuizRepository struct {
	Db *gorm.DB
}

func NewQuizRepository() *QuizRepository {
	return &QuizRepository{
		Db: configs.GetDB(),
	}
}

func (qr *QuizRepository) Create(quiz models.Quiz) error {
	tx := qr.Db.Begin()
	err := tx.Create(&quiz).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (qr *QuizRepository) All() ([]models.Quiz, error) {
	var quizzes []models.Quiz
	err := qr.Db.Find(&quizzes).Error
	return quizzes, err
}

func (qr *QuizRepository) FindID(id uint) (models.Quiz, error) {
	var quiz models.Quiz
	err := qr.Db.Where("id = ?", id).First(&quiz).Error
	return quiz, err
}

func (qr *QuizRepository) Free() ([]models.Quiz, error) {
	var freeQuizzes []models.Quiz
	//SELECT * FROM quizzes WHERE free = true
	err := qr.Db.Where("free = ?", true).Find(&freeQuizzes).Error
	return freeQuizzes, err
}

func (qr *QuizRepository) Cheapest() ([]models.Quiz, error) {
	var cheapQuizzes []models.Quiz
	//SELECT * FROM quizzess ORDER BY price ASC
	err := qr.Db.Where("free = ?", false).Find(&cheapQuizzes).Order("price ASC").Error
	return cheapQuizzes, err
}

func (qr *QuizRepository) FindTitle(title string) ([]models.Quiz, error) {
	var quiz []models.Quiz
	err := qr.Db.Where("title = ?", "%"+title+"%").Find(&quiz).Error
	return quiz, err
}

func (qr *QuizRepository) FindTopic(topic string) ([]models.Quiz, error) {
	var quiz []models.Quiz
	err := qr.Db.Where("topic = ?", "%"+topic+"%").Find(&quiz).Error
	return quiz, err
}

func (qr *QuizRepository) Update(quiz models.Quiz) error {
	tx := qr.Db.Begin()
	//UPDATE quizzes SET VALUE({quiz}) WHERE id = {quiz->id}
	err := tx.Model(&models.Quiz{}).Where("id = ?", quiz.ID).Updates(&quiz).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

/*set new Discount value to : <id> <discount>*/
func (qr *QuizRepository) SetDiscount(id uint, disc uint8) error {
	tx := qr.Db.Begin()
	err := tx.Model(&models.Quiz{}).Where("id = ? AND free = ?", id, false).Update("disc", disc).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (qr *QuizRepository) Delete(id uint) error {
	tx := qr.Db.Begin()
	err := tx.Model(&models.Quiz{}).Where("id = ?", id).Update("deleted_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

package service

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/dto"
	"BagasA11/GSC-quizHealthEdu-BE/api/models"
	"BagasA11/GSC-quizHealthEdu-BE/api/repository"
)

type QuestionService struct {
	repository *repository.QuestionRepository
}

func NewQuestionService() *QuestionService {
	return &QuestionService{
		repository: repository.NewQuestionRepository(),
	}
}

func (qst *QuestionService) Create(req *dto.Question) error {
	q := models.Question{
		Question: req.Question,
		Answer:   req.Answer,
	}
	err := qst.repository.Create(q)
	return err
}

func (qst *QuestionService) ReferToQuiz(quizID uint) ([]models.Question, error) {
	q, err := qst.repository.ReferToQuiz(quizID)
	return q, err
}

func (qst *QuestionService) Updates(id uint, req *dto.Question) error {
	q := models.Question{
		ID:       id,
		Question: req.Question,
		Answer:   req.Answer,
	}
	err := qst.repository.Updates(q)
	return err
}

func (qst *QuestionService) Delete(id uint) error {
	err := qst.repository.Delete(id)
	return err
}

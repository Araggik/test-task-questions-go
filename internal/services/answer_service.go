package services

import (
	"errors"

	"github.com/Araggik/test-task-questions-go/internal/models"
	"github.com/Araggik/test-task-questions-go/internal/repositories"
)

type AnswerService interface {
	GetAnswer(id int) (*models.Answer, error)
	CreateAnswer(req models.CreateAnswerRequest) (*models.Answer, error)
	DeleteAnswer(id int) error
}

type answerService struct {
	repo repositories.AnswerRepository
}

func NewAnswerService(repo repositories.AnswerRepository) AnswerService {
	return &answerService{repo: repo}
}

func (s *answerService) GetAnswer(id int) (*models.Answer, error) {
	var err error
	var answer *models.Answer

	if id <= 0 {
		err = errors.New("id должен быть больше нуля")
	}

	if err == nil {
		answer, err = s.repo.GetByID(id)
	}

	return answer, err
}

func (s *answerService) CreateAnswer(req models.CreateAnswerRequest) (*models.Answer, error) {
	answer := &models.Answer{
		Text:       req.Text,
		QuestionID: req.QuestionID,
		UserID:     req.UserID,
	}

	err := s.repo.Create(answer)

	return answer, err
}

func (s *answerService) DeleteAnswer(id int) error {
	var err error

	if id <= 0 {
		err = errors.New("id должен быть больше нуля")
	}

	if err == nil {
		err = s.repo.Delete(id)
	}

	return err
}

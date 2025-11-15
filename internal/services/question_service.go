package services

import (
	"errors"

	"github.com/Araggik/test-task-questions-go/internal/models"
	"github.com/Araggik/test-task-questions-go/internal/repositories"
)

type QuestionService interface {
	GetQuestion(id int) (*models.Question, error)
	CreateQuestion(req models.CreateQuestionRequest) (*models.Question, error)
}

type questionService struct {
	repo repositories.QuestionRepository
}

func NewQuestionService(repo repositories.QuestionRepository) QuestionService {
	return &questionService{repo: repo}
}

func (s *questionService) GetQuestion(id int) (*models.Question, error) {
	var err error
	var question *models.Question

	if id <= 0 {
		err = errors.New("id должен быть больше нуля")
	}

	if err == nil {
		question, err = s.repo.GetByID(id)
	}

	return question, err
}

func (s *questionService) CreateQuestion(req models.CreateQuestionRequest) (*models.Question, error) {
	question := &models.Question{
		Text: req.Text,
	}

	err := s.repo.Create(question)

	return question, err
}

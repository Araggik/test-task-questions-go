package services

import (
	"errors"

	"github.com/Araggik/test-task-questions-go/internal/models"
	"github.com/Araggik/test-task-questions-go/internal/repositories"
)

type QuestionService interface {
	GetQuestion(id int) (*models.Question, error)
	CreateQuestion(req models.CreateQuestionRequest) (*models.Question, error)
	GetAllQuestions() ([]models.Question, error)
	DeleteQuestion(id int) error
}

type questionService struct {
	questionRepo repositories.QuestionRepository
}

func NewQuestionService(questionRepo repositories.QuestionRepository) QuestionService {
	return &questionService{questionRepo: questionRepo}
}

func (s *questionService) GetQuestion(id int) (*models.Question, error) {
	var err error
	var question *models.Question

	if id <= 0 {
		err = errors.New("id должен быть больше нуля")
	}

	if err == nil {
		question, err = s.questionRepo.GetByID(id)
	}

	return question, err
}

func (s *questionService) CreateQuestion(req models.CreateQuestionRequest) (*models.Question, error) {
	question := &models.Question{
		Text: req.Text,
	}

	err := s.questionRepo.Create(question)

	return question, err
}

func (s *questionService) GetAllQuestions() ([]models.Question, error) {
	questions, err := s.questionRepo.GetAll()

	return questions, err
}

func (s *questionService) DeleteQuestion(id int) error {
	var err error

	if id <= 0 {
		err = errors.New("id должен быть больше нуля")
		return err
	}

	//Проверяем, что есть вопрос с указанным id
	_, err = s.questionRepo.GetByID(id)
	if err != nil {
		return err
	}

	//Каскадное удаление на уровне бд
	err = s.questionRepo.Delete(id)

	return err
}

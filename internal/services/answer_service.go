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
	answerRepo   repositories.AnswerRepository
	questionRepo repositories.QuestionRepository
}

func NewAnswerService(answerRepo repositories.AnswerRepository, questionRepo repositories.QuestionRepository) AnswerService {
	return &answerService{answerRepo: answerRepo, questionRepo: questionRepo}
}

func (s *answerService) GetAnswer(id int) (*models.Answer, error) {
	var err error
	var answer *models.Answer

	if id <= 0 {
		err = errors.New("id должен быть больше нуля")
	}

	if err == nil {
		answer, err = s.answerRepo.GetByID(id)
	}

	return answer, err
}

func (s *answerService) CreateAnswer(req models.CreateAnswerRequest) (*models.Answer, error) {
	//Проверяем, что есть указанный вопрос
	_, err := s.questionRepo.GetByID(req.QuestionID)
	if err != nil {
		return nil, err
	}

	answer := &models.Answer{
		Text:       req.Text,
		QuestionID: req.QuestionID,
		UserID:     req.UserID,
	}

	err = s.answerRepo.Create(answer)

	return answer, err
}

func (s *answerService) DeleteAnswer(id int) error {
	var err error

	if id <= 0 {
		err = errors.New("id должен быть больше нуля")
		return err
	}

	//Проверяем, что есть ответ с указанным id
	_, err = s.answerRepo.GetByID(id)
	if err != nil {
		return err
	}

	err = s.answerRepo.Delete(id)

	return err
}

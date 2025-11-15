package services

import "github.com/Araggik/test-task-questions-go/internal/repositories"

type AnswerService interface {
}

type answerService struct {
	repo repositories.AnswerRepository
}

func NewAnswerService(repo repositories.AnswerRepository) AnswerService {
	return &answerService{repo: repo}
}

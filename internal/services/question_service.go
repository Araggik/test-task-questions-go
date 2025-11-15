package services

import "github.com/Araggik/test-task-questions-go/internal/repositories"

type QuestionService interface {
}

type questionService struct {
	repo repositories.QuestionRepository
}

func NewQuestionService(repo repositories.QuestionRepository) QuestionService {
	return &questionService{repo: repo}
}

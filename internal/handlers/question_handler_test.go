package handlers

import (
	"errors"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Araggik/test-task-questions-go/internal/models"
	"github.com/Araggik/test-task-questions-go/internal/services"
	"github.com/stretchr/testify/assert"
)

type testQuestionService struct {
}

func NewTestQuestionService() services.QuestionService {
	return &testQuestionService{}
}

func (s *testQuestionService) GetQuestion(id int) (*models.Question, error) {
	if id == 1 {
		return &models.Question{
			ID:        1,
			Text:      "Test",
			CreatedAt: time.Date(2025, time.November, 15, 0, 0, 0, 0, time.UTC),
		}, nil
	} else {
		return nil, errors.New("Нет такого вопроса")
	}
}

func (s *testQuestionService) CreateQuestion(req models.CreateQuestionRequest) (*models.Question, error) {
	question := &models.Question{
		ID:        2,
		Text:      req.Text,
		CreatedAt: time.Now(),
	}

	return question, nil
}

func (s *testQuestionService) GetAllQuestions() ([]models.Question, error) {
	return []models.Question{}, nil
}

func (s *testQuestionService) DeleteQuestion(id int) error {
	return nil
}

func TestGetQuestion(t *testing.T) {
	testService := NewTestQuestionService()
	questionHandler := NewQuestionHandler(testService)
	handler := questionHandler.GetQuestion

	req := httptest.NewRequest("GET", "http://example.com/questions/1", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, string(body), `{"id":1,"text":"Test","created_at":"2025-11-15T00:00:00Z"}`)
}

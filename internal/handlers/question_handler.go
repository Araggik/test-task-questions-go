package handlers

import (
	"io"
	"net/http"

	"github.com/Araggik/test-task-questions-go/internal/services"
)

type QuestionHandler struct {
	service services.QuestionService
}

func NewQuestionHandler(service services.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: service}
}

func (h *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GetQuestion")
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "CreateQuestion")
}

func (h *QuestionHandler) GetAllQuestion(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GetAllQuestion")
}

func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "DeleteQuestion")
}

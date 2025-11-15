package handlers

import (
	"io"
	"net/http"

	"github.com/Araggik/test-task-questions-go/internal/services"
)

type AnswerHandler struct {
	service services.AnswerService
}

func NewAnswerHandler(service services.AnswerService) *AnswerHandler {
	return &AnswerHandler{service: service}
}

func (h *AnswerHandler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GetAnswer")
}

func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "CreateAnswer")
}

func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "DeleteAnswer")
}

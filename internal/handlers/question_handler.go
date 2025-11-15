package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Araggik/test-task-questions-go/internal/services"
)

type QuestionHandler struct {
	service services.QuestionService
}

func NewQuestionHandler(service services.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: service}
}

func (h *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	rawId := parts[2]

	id, err := strconv.Atoi(rawId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Некорректный id в запросе"))
		return
	}

	question, err := h.service.GetQuestion(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(*question)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {

}

func (h *QuestionHandler) GetAllQuestion(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GetAllQuestion")
}

func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "DeleteQuestion")
}

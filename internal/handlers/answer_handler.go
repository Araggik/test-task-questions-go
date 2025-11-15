package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Araggik/test-task-questions-go/internal/models"
	"github.com/Araggik/test-task-questions-go/internal/services"
)

type AnswerHandler struct {
	service services.AnswerService
}

func NewAnswerHandler(service services.AnswerService) *AnswerHandler {
	return &AnswerHandler{service: service}
}

func (h *AnswerHandler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	rawId := parts[2]

	id, err := strconv.Atoi(rawId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Некорректный id в запросе"))
		return
	}

	answer, err := h.service.GetAnswer(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(*answer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAnswerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Правильный формат json: {"text": "Текст вопроса", "user_id": "38473b7f-c618-4f80-ae37-e22fefc01aa8", "question_id": 1})`))
		return
	}

	answer, err := h.service.CreateAnswer(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(*answer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	rawId := parts[2]

	id, err := strconv.Atoi(rawId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Некорректный id в запросе"))
		return
	}

	err = h.service.DeleteAnswer(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

package repositories

import (
	"github.com/Araggik/test-task-questions-go/internal/models"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	Create(question *models.Answer) error
	GetByID(id int) (*models.Answer, error)
	Delete(id int) error
}

type answerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerRepository{db: db}
}

func (r *answerRepository) Create(question *models.Answer) error {
	return r.db.Create(question).Error
}

func (r *answerRepository) GetByID(id int) (*models.Answer, error) {
	question := new(models.Answer)

	err := r.db.First(question, id).Error

	return question, err
}

func (r *answerRepository) Delete(id int) error {
	return r.db.Delete(&models.Answer{}, id).Error
}

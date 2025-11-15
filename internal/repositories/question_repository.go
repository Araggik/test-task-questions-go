package repositories

import (
	"github.com/Araggik/test-task-questions-go/internal/models"
	"gorm.io/gorm"
)

type QuestionRepository interface {
	Create(question *models.Question) error
	GetByID(id int) (*models.Question, error)
	GetAll() ([]models.Question, error)
	Delete(id int) error
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) Create(question *models.Question) error {
	return r.db.Create(question).Error
}

func (r *questionRepository) GetByID(id int) (*models.Question, error) {
	question := new(models.Question)

	err := r.db.First(question, id).Error

	return question, err
}

func (r *questionRepository) GetAll() ([]models.Question, error) {
	var questions []models.Question

	err := r.db.Find(&questions).Error

	return questions, err
}

func (r *questionRepository) Delete(id int) error {
	return r.db.Delete(&models.Question{}, id).Error
}

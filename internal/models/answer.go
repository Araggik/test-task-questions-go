package models

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	ID         int
	QuestionID int
	UserID     string
	Text       string
	CreatedAt  time.Time
}

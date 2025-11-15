package models

import (
	"time"
)

type Answer struct {
	ID         int       `json:"id"  gorm:"primaryKey;autoIncrement"`
	QuestionID int       `json:"question_id" gorm:"not null;index"`
	UserID     string    `json:"user_id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}

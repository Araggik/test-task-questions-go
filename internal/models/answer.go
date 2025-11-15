package models

import (
	"time"
)

type Answer struct {
	ID         int       `json:"id"  gorm:"primaryKey;autoIncrement"`
	QuestionID int       `json:"question_id" gorm:"not null;index"`
	UserID     string    `json:"user_id"`
	Text       string    `json:"text" gorm:"column:txt"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateAnswerRequest struct {
	QuestionID int    `json:"question_id"`
	UserID     string `json:"user_id"`
	Text       string `json:"text"`
}

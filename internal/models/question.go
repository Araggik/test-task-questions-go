package models

import (
	"time"
)

type Question struct {
	ID        int       `json:"id"  gorm:"primaryKey;autoIncrement"`
	Text      string    `json:"text" gorm:"column:txt"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateQuestionRequest struct {
	Text string `json:"text"`
}

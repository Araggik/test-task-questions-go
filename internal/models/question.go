package models

import (
	"time"
)

type Question struct {
	ID        int       `json:"id"  gorm:"primaryKey;autoIncrement"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

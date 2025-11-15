package models

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID        int
	Text      string
	CreatedAt time.Time
}

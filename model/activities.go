package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID          int64 `gorm:"primarykey"`
	Description string

	UserID int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

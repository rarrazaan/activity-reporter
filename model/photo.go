package model

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID      int64 `gorm:"primarykey"`
	Caption string

	UserID int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

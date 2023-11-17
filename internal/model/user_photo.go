package model

import (
	"time"

	"gorm.io/gorm"
)

type UserPhoto struct {
	ID int64 `gorm:"primaryKey"`

	UserID  int64
	PhotoID int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

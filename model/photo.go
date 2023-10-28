package model

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID      int64  `gorm:"primarykey"`
	Image   string `gorm:"not null;type:varchar"`
	Caption string

	UserID int64
	Likers []User `gorm:"many2many:user_photos;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

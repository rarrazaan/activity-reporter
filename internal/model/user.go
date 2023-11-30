package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int64  `gorm:"primarykey"`
	Name            string `gorm:"type:varchar;not null"`
	Username        string `gorm:"type:varchar;not null;index:idx_username,unique"`
	Email           string `gorm:"type:varchar;not null;index:idx_email,unique"`
	EmailIsVerified bool   `gorm:"not null;default:FALSE"`
	Password        string `gorm:"type:varchar; not null"`

	Followers  []User     `gorm:"foreignKey:ID"`
	Activities []Activity `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

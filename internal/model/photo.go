package model

import (
	"time"
)

type Photo struct {
	ID        string    `bson:"id"`
	ImageUrl  string    `bson:"image_url"`
	Caption   string    `bson:"caption"`
	UserID    int64     `bson:"user_id"`
	Username  string    `bson:"user_name"`
	Likers    []liker   `bson:"likers"`
	CreatedAt time.Time `bson:"created_at"`
}

type liker struct {
	UserID   int64  `bson:"user_id"`
	Username string `bson:"user_name"`
}

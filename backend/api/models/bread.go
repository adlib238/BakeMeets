package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `json:"content" gorm:"not null"`
	BreadID uint   `json:"bread_id"`
	UserID	uint   `json:"user_id"`
}

type Bread struct {
	gorm.Model
	Name          string `json:"name" gorm:"not null"`
	Description   string `json:"description"`
	FavoriteCount int    `json:"favorite_count" gorm:"default:0"`
	CommentCount  int    `json:"comment_count" gorm:"default:0"`
	ImageURL      string `json:"image_url"`
}

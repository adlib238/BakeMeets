package models

import (
	"gorm.io/gorm"
)

type Bread struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
}

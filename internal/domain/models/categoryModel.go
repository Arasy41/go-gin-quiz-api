package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"not null" json:"id"`
	Name string `gorm:"not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type CategoryList struct {
	ID   uint   `gorm:"not null" json:"id"`
	Name string `gorm:"not null" json:"name"`
}
package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `gorm:"not null" json:"id"`
	Name      string         `gorm:"not null;unique" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	User      UserList       `json:"user,omitempty"`
}

type RoleList struct {
	ID   uint   `gorm:"not null" json:"id"`
	Name string `gorm:"not null;unique" json:"name"`
}

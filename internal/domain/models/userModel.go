package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"not null;unique" json:"username"`
	Email     string         `gorm:"not null;unique" json:"email"`
	Password  string         `gorm:"not null" json:"password"`
	RoleID    uint           `gorm:"not null" json:"role_id"`
	Role      Role           `gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserList struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleName string `json:"role_name"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

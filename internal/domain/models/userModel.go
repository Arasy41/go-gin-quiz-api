package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"not null;unique" json:"username"`
	Email     string `gorm:"not null;unique" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	RoleID    uint   `gorm:"not null" json:"role_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
	Role      Role      `gorm:"foreignKey:RoleID"`
}

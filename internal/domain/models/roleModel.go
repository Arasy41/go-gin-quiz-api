package models

import "time"

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null;unique" json:"name"`

	Users []User `gorm:"foreignKey:RoleID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
}

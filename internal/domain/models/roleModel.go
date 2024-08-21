package models

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null;unique" json:"name"`
}
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quiz struct {
	ID          uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string        `gorm:"type:varchar(255);not null" json:"title"`
	Description string        `gorm:"type:text" json:"description"`
	CategoryID  uint          `gorm:"not null" json:"category_id"`
	Difficulty  string        `gorm:"not null" json:"difficulty"`
	Duration    time.Duration `gorm:"not null" json:"duration"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Questions   []Question    `gorm:"foreignKey:QuizID" json:"questions"`
}

func (quiz *Quiz) BeforeCreate(tx *gorm.DB) (err error) {
	quiz.ID = uuid.New()
	return nil
}

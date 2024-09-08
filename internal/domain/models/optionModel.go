package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Option struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	QuestionID uuid.UUID `json:"question_id"`
	Text       string    `json:"text"`
}

func (option *Option) BeforeCreate(tx *gorm.DB) (err error) {
	option.ID = uuid.New()
	return nil
}

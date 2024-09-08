package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
    ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
    QuizID   uuid.UUID `json:"quiz_id"`
    Text     string    `json:"text"`
    Options  []Option  `json:"options"`
    AnswerID uuid.UUID `json:"answer_id"`  // ID jawaban yang benar
}

func (question *Question) BeforeCreate(tx *gorm.DB) (err error) {
    question.ID = uuid.New()
    return nil
}
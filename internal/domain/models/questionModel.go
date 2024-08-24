package models

import (
	"github.com/google/uuid"
)

type Question struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	QuizID  uuid.UUID `json:"quiz_id"`
	Text    string    `json:"text"`
	Options []Option  `json:"options"` // Multiple options for each question
	Answer  string    `json:"answer"`  // Correct answer
}

type Option struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	QuestionID uuid.UUID `json:"question_id"`
	Text       string    `json:"text"`
}

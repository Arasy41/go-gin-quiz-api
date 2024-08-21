package models

import "github.com/google/uuid"

type Question struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	QuizID      uuid.UUID `gorm:"type:uuid" json:"quiz_id"`
	Question    string    `gorm:"type:text" json:"question"`
	Answer      string    `gorm:"type:text" json:"answer"`
	Difficulty  string    `gorm:"not null" json:"difficulty"`
	Explanation string    `gorm:"type:text" json:"explanation"`
}

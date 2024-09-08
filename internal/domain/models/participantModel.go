package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Participant struct {
    ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    QuizID   uuid.UUID `json:"quiz_id"`
    UserID   uint      `json:"user_id"`
    Score    int       `json:"score"`
    Finished bool      `json:"finished"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (participant *Participant) BeforeCreate(tx *gorm.DB) (err error) {
    participant.ID = uuid.New()
    return nil
}
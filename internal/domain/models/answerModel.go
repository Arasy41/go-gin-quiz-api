package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Answer struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ParticipantID uuid.UUID `json:"participant_id"`
	QuestionID    uuid.UUID `json:"question_id"`
	OptionID      uuid.UUID `json:"option_id"`
	Correct       bool      `json:"correct"`
}

func (answer *Answer) BeforeCreate(tx *gorm.DB) (err error) {
	answer.ID = uuid.New()
	return nil
}

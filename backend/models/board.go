package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Board struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Title     string    `json:"title" binding:"required"`
	Tasks     []Task    `json:"tasks,omitempty"` // Panoya ait görevleri yüklemek için
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (board *Board) BeforeCreate(tx *gorm.DB) (err error) {
	board.ID = uuid.New()
	return
}

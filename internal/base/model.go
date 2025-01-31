package base

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}

func NewUUID() uuid.UUID {
	return uuid.New()
}

func (e *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuid.Nil {
		e.ID = NewUUID()
	}

	if e.CreatedAt.IsZero() {
		e.CreatedAt = time.Now()
	}

	if e.UpdatedAt.IsZero() {
		e.UpdatedAt = time.Now()
	}

	return nil
}

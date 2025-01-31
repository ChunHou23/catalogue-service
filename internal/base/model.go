package base

import (
	"time"

	"github.com/google/uuid"
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

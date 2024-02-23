package models

import (
	"github.com/nexodus-io/nexodus/internal/database/datatype"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StringArray = datatype.StringArray

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id" example:"aa22666c-0f57-45cb-a449-16efecc04f2e"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate populates the ID (if not set)
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return nil
}

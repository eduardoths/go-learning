package migrations

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
}

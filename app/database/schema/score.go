package schema

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	User   User
	UserID string `gorm:"not null"`
	Score  int
}

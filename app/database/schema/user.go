package schema

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   string
	Name string
}

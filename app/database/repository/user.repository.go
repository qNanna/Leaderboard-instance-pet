package repository

import (
	"gorm.io/gorm"

	"main/app/database"
	"main/app/database/schema"
)

type UserRepository struct{}

func (u UserRepository) SaveUser(uid string) *gorm.DB {
	user := schema.User{}

	return database.GetConnection().FirstOrCreate(&user, schema.User{ID: uid})
}

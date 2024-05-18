package database

import (
	"errors"
	"fmt"
	"os"

	"main/app/database/schema"

	"github.com/gofor-little/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Init() {
	path := env.Get("DATABASE_PATH", "database.db")
	initDatabase(path)

	var err error
	connection, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	connection.AutoMigrate(&schema.User{}, &schema.Score{})
}

func createDatabase(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Database file not created")
		return
	}
	defer file.Close()
}

func initDatabase(filePath string) {
	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Database doesn't exist")
		createDatabase(filePath)
	}
}

func GetConnection() *gorm.DB {
	return connection
}

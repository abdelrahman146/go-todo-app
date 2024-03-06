package providers

import (
	"log"
	"todo-app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database:", err)
	}

	log.Println("Database connected")
	// Migrate the schema, you can add other model migrations here as well
	DB.AutoMigrate(&models.Todo{})
}

func GetDB() *gorm.DB {
	return DB
}

package config

import (
	"fmt"
	"golang-todo-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", env.DB_HOST, env.DB_USER, env.DB_PASSWORD, env.DB_NAME, env.DB_PORT, env.DB_SSL_MODE, env.DB_TIMEZONE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(models.Todo{})
	if err != nil {
		log.Fatalln(err)
	}
}

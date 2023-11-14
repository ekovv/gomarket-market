package db

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDB(config config.Config) *gorm.DB {
	url := config.DB
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}

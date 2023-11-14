package main

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/internal/db"
	model2 "go-rest-api/internal/model"
)

func main() {
	cnf := config.New()
	dbConn := db.NewDB(cnf)
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model2.User{})
}

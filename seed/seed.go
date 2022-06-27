package main

import (
	"go-rest-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(localhost:3307)/go_api?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Username: "Admin", Email: "admin@tester.com", Password: "admin"})
	db.Create(&model.User{Username: "Test", Email: "tester@tester.com", Password: "test"})

}

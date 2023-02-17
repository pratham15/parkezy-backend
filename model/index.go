package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// connect to db
	dsn := "root:12345678@tcp(127.0.0.1:3306)/parkezy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	err = db.AutoMigrate(&User{}, &Parking{}, &Booking{})
	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("Database connected")
	DB = db
}

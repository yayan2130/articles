package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB variable to be used in controllers
var DB *gorm.DB

func ConnectDatabase() {
	fmt.Println("🔄 Connecting to database...")

	dsn := "root:@tcp(127.0.0.1:3306)/article?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Database connection failed: ", err)
	}

	DB = database
	fmt.Println("✅ Database connected successfully")
}
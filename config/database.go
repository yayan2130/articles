package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/article?charset=utf8mb4&parseTime=True&loc=Local" // Sesuaikan dengan XAMPP kamu
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	DB = db
	fmt.Println("Database berhasil terhubung! 🚀")
}

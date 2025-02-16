package main

import (
	"article/config"
	"article/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode) // Ubah ke gin.ReleaseMode di produksi

	// Koneksi database
	config.ConnectDatabase()

	// Inisialisasi router
	r := gin.Default()

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Next.js
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Daftarkan routes
	routes.SetupRouter(r)

	// Jalankan server
	err := r.Run(":8080")
	if err != nil {
		panic("Server gagal berjalan: " + err.Error())
	}
}

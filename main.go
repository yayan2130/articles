package main

import (
	"github.com/gin-gonic/gin"

	"article/config"
	"article/routes"
)

func main() {
	// Pastikan mode diatur agar tidak ada warning (opsional)
	gin.SetMode(gin.DebugMode) // atau gin.ReleaseMode jika di produksi
	
	// Panggil koneksi database sebelum router dijalankan
	config.ConnectDatabase()

	// Panggil router dari routes.go
	r := routes.SetupRouter()

	// Jalankan server di port 8080
	err := r.Run(":8080")
	if err != nil {
		panic("Server gagal berjalan: " + err.Error())
	}
}
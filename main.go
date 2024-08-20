package main

import (
	"gin-rest-api/handlers"
	"gin-rest-api/middleware"
	"gin-rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
    // Inisialisasi koneksi database
    models.InitDB("system@JWI88:system@JWI88@tcp(localhost:3306)/ginapi")

    r := gin.Default()

    // Rute untuk pendaftaran dan login
    r.POST("/api/register", handlers.Register)
    r.POST("/api/login", handlers.Login)

    // Rute yang memerlukan autentikasi
    authorized := r.Group("/")
    authorized.Use(middleware.AuthMiddleware())
    {
        authorized.GET("/api/items", handlers.GetItems)
        authorized.GET("/api/items/:id", handlers.GetItem)
        authorized.POST("/api/items", handlers.CreateItem)
        authorized.PUT("/api/items/:id", handlers.UpdateItem)
        authorized.DELETE("/api/items/:id", handlers.DeleteItem)
    }

    r.Run(":8000") // Menjalankan server di port 8000
}

package main

import (
    "github.com/gin-gonic/gin"
    "go-crud-api/database"
    "go-crud-api/models"
    "go-crud-api/routes"
	"go-crud-api/middleware"
	"github.com/gin-contrib/cors"
)

func main() {
    // 1. Koneksi database
    database.Connect()
    database.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Log{}, &models.Activity{}) // otomatis bikin tabel user dan post

    // 2. Init Gin
    r := gin.Default()

    // middleware CORS
	r.Use(cors.Default())

    // 3. Routes
    routes.AuthRoutes(r)
    
    // Protected routes
    r.Use(middleware.AuthMiddleware())
    routes.UserRoutes(r) // load user route
    routes.PostRoutes(r) // load post route
    routes.ActivityRoutes(r) // load activity route
    routes.DashboardRoutes(r) // load dashboard route

    // 4. Run server
    r.Run(":8081") // default localhost:8080
}

package main

import (
    "github.com/gin-gonic/gin"
    "go-crud-api/database"
    "go-crud-api/models"
    "go-crud-api/routes"
)

func main() {
    // 1. Koneksi database
    database.Connect()
    database.DB.AutoMigrate(&models.User{}) // otomatis bikin tabel user

    // 2. Init Gin
    r := gin.Default()

    // 3. Routes
    r.GET("/users", routes.GetUsers)
    r.POST("/users", routes.CreateUser)
    r.GET("/users/:id", routes.GetUser)
    r.PUT("/users/:id", routes.UpdateUser)
    r.DELETE("/users/:id", routes.DeleteUser)

    // 4. Run server
    r.Run(":8080") // default localhost:8080
}

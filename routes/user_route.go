package routes

import (
    "github.com/gin-gonic/gin"
    "go-crud-api/controllers"
)

func UserRoutes(r *gin.Engine) {
    r.GET("/users", controllers.GetUsers)
    r.POST("/users", controllers.CreateUser)
    r.GET("/users/:id", controllers.GetUser)
    r.PUT("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)
}
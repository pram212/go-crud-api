package routes

import (
    "github.com/gin-gonic/gin"
    "go-crud-api/controllers"
)

func AuthRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)
}
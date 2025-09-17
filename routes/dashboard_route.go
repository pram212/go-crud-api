package routes

import (
    "github.com/gin-gonic/gin"
    "go-crud-api/controllers"
)

func DashboardRoutes(r *gin.Engine) {
    r.GET("/dashboard", controllers.Dashboard)
}

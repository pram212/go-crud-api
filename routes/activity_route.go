package routes

import (
	"go-crud-api/controllers"
	"github.com/gin-gonic/gin"
)

func ActivityRoutes(router *gin.Engine) {
	router.GET("/activities", controllers.GetAllActivities)
	router.GET("/activities/:id", controllers.GetUserActivities)
}

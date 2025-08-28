package routes

import (
	"github.com/gin-gonic/gin"
	"go-crud-api/controllers"
)

func PostRoutes(r *gin.Engine) {
	r.GET("/posts", controllers.GetPosts)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
}
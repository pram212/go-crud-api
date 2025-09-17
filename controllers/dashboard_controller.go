package controllers

import (
    "go-crud-api/database"
    "go-crud-api/models"
    "go-crud-api/utils"
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
    userChan := make(chan []models.User)
    postChan := make(chan []models.Post)
	activitiesChan := make(chan []models.Activity)

    // ambil users
    go func() {
        var users []models.User
        database.DB.Find(&users)
        userChan <- users
    }()

    // ambil posts
    go func() {
        var posts []models.Post
        database.DB.Find(&posts)
        postChan <- posts
    }()

	// ambil activities
	go func() {
		var activities []models.Activity
		database.DB.Find(&activities)
		activitiesChan <- activities
	}()

    users := <-userChan
    posts := <-postChan
	activities := <-activitiesChan

	// return response
    utils.SuccessResponse(c, "Dashboard data retrieved successfully", map[string]interface{}{
        "posts": posts,
		"users": users,
        "activities": activities,
    })

}

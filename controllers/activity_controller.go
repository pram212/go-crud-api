package controllers

import (
	"go-crud-api/database"
	"go-crud-api/models"
	"github.com/gin-gonic/gin"
	"go-crud-api/utils"
)

// get activities by user id
func GetUserActivities(c *gin.Context) {
	userID := c.Param("id")
	var activities []models.Activity
	database.DB.Where("user_id = ?", userID).Find(&activities)
	utils.SuccessResponse(c, "User activities retrieved successfully", activities)
}

// get all activities
func GetAllActivities(c *gin.Context) {
	var activities []models.Activity
	database.DB.Find(&activities)
	utils.SuccessResponse(c, "All activities retrieved successfully", activities)
}
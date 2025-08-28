package controllers

import (
	"github.com/gin-gonic/gin"
	"go-crud-api/database"
	"go-crud-api/models"
	"net/http"
	"go-crud-api/utils"
)

// GET /posts
func GetPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Preload("User").Find(&posts) // ikut ambil data user
	utils.SuccessResponse(c, "Posts retrieved successfully", posts)
}

// POST /posts
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&post).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, "Post created successfully", post)
}

// GET /posts/:id
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.Preload("User").First(&post, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}
	utils.SuccessResponse(c, "Post retrieved successfully", post)
}

// PUT /posts/:id
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	database.DB.Save(&post)
	utils.SuccessResponse(c, "Post updated successfully", post)
}

// DELETE /posts/:id
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Post{}, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.SuccessResponse(c, "Post deleted successfully", nil)
}
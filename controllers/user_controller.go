package controllers

import (
	"go-crud-api/database"
	"go-crud-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"go-crud-api/utils"
)

var validate = validator.New()

// GET /users
func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Preload("Posts").Find(&users)
    utils.SuccessResponse(c, "Users retrieved successfully", users)
}

// CreateUser controller
func CreateUser(c *gin.Context) {
	var user models.User

	// Ambil data dari request body
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Validasi field
	if err := validate.Struct(user); err != nil {
		errors := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" is invalid")
		}
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", gin.H{"validation_errors": errors})
		return
	}

	// Hash password sebelum simpan
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, "User created successfully", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

// GET /users/:id
func GetUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "User not found")
        return
    }
    utils.SuccessResponse(c, "User retrieved successfully", user)
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "User not found")
        return
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    database.DB.Save(&user)
    utils.SuccessResponse(c, "User updated successfully", user)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := database.DB.Delete(&user, id).Error; err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    utils.SuccessResponse(c, "User deleted successfully", nil)
}

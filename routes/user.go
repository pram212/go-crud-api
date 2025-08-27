package routes

import (
    "github.com/gin-gonic/gin"
    "go-crud-api/database"
    "go-crud-api/models"
    "net/http"
    "github.com/go-playground/validator/v10"
)

var validate = validator.New()

// GET /users
func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

// POST /users
func CreateUser(c *gin.Context) {
    var user models.User
    // Ambil data dari request body
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validasi field sesuai tag di model
    if err := validate.Struct(user); err != nil {
        errors := []string{}
        for _, err := range err.(validator.ValidationErrors) {
            var msg string
            switch err.Tag() {
            case "required":
                msg = err.Field() + " is required"
            case "email":
                msg = err.Field() + " must be a valid email address"
            case "min":
                msg = err.Field() + " must be at least " + err.Param() + " characters"
            default:
                msg = "Invalid value for " + err.Field()
            }
            errors = append(errors, msg)
        }
        c.JSON(http.StatusBadRequest, gin.H{"validation_errors": errors})
        return
    }


    database.DB.Create(&user)
    c.JSON(http.StatusOK, user)
}

// GET /users/:id
func GetUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Save(&user)
    c.JSON(http.StatusOK, user)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := database.DB.Delete(&user, id).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

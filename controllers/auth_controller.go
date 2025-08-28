package controllers

import (
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
    "go-crud-api/database"
    "go-crud-api/models"
    "go-crud-api/utils"
    "net/http"
    "os"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type LoginResponse struct {
    Token string      `json:"token"`
    User  UserResponse `json:"user"`
}

type UserResponse struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
        return
    }

    var user models.User
    
    if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        utils.ErrorResponse(c, http.StatusUnauthorized, "User not found")
        return
    }

    // Bandingkan password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password")
        return
    }

    // generate token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": user.ID,
        "email": user.Email,
    })

    // sign token
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
        return
    }

    res := LoginResponse{
        Token: tokenString,
        User: UserResponse{
            ID:    user.ID,
            Name:  user.Name,
            Email: user.Email,
        },
    }

    utils.SuccessResponse(c, "Login successful", res)
}
package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret_key_123") // nanti sebaiknya simpan di .env

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        // Format: "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
            c.Abort()
            return
        }

        tokenStr := parts[1]

        // Parse token
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        c.Next() // lanjut ke handler berikutnya
    }
}
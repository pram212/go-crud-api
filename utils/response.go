package utils

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
    res := Response{
        Status:  "success",
        Message: message,
        Data:    data,
    }
    c.JSON(http.StatusOK, res)
}

func ErrorResponse(c *gin.Context, code int, message string, data ...interface{}) {
    response := gin.H{
        "status":  "error",
        "message": message,
    }

    if len(data) > 0 {
        response["data"] = data[0] // ambil hanya 1 kalau ada
    }

    c.JSON(code, response)
}


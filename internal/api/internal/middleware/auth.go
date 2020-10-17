package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func IsAuthorised(c *gin.Context) {
    _, exists := c.Get("user")

    if !exists {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }
}

package auth

import (
    "github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
    auth := r.Group("/auth")
    {
        auth.POST("/signup", signup)
        auth.POST("/signin", signin)
        auth.GET("/check", check)
    }
}

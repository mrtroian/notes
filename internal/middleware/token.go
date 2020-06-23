package middleware

import (
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/mrtroian/notes/internal/common"
    "github.com/mrtroian/notes/internal/database/models"
    "github.com/mrtroian/notes/internal/token"
)

func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.User

        tokenString, err := c.Cookie("token")

        if err != nil {
            authorization := c.Request.Header.Get("Authorization")

            if len(authorization) == 0 {
                c.Next()
                return
            }
            sp := strings.Split(authorization, "Bearer ")

            if len(sp) < 1 {
                c.Next()
                return
            }
            tokenString = sp[1]
        }
        tokenData, err := token.Validate(tokenString)

        if err != nil {
            c.Next()
            return
        }

        user.Read(tokenData["user"].(common.JSON))

        c.Set("user", user)
        c.Set("token_expire", tokenData["exp"])
        c.Next()
    }
}

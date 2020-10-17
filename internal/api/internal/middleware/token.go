package middleware

import (
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/mrtroian/notes/internal/token"
    "github.com/mrtroian/notes/internal/user"
)

func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString, err := c.Cookie("token")

        if err != nil {
            // in case user cleaned up all the cookies
            // but he stil has set auth header with valid token
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

        u, err := user.ReadJSON(tokenData["user"].(map[string]interface{}))

        if err != nil {
            c.Next()
            return
        }

        // unauthorised users will not appear in gin.Context
        c.Set("user", u)
        c.Set("token_expire", tokenData["exp"])
        c.Next()
    }
}

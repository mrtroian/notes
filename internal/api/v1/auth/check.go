package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrtroian/notes/internal/token"
	"github.com/mrtroian/notes/internal/user"
)

const weekTime int = 60 * 60 * 24 * 7

func check(c *gin.Context) {
	userRaw, ok := c.Get("user")

	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user := userRaw.(*user.User)
	t, err := token.Generate(user)

	if err != nil {
		log.Println(err)
	}

	c.SetCookie("token", t, weekTime, "/", "", false, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": t,
	})
}

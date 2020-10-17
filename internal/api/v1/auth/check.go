package auth

import (
	"log"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/mrtroian/notes/internal/common"
	"github.com/mrtroian/notes/internal/token"
)

const weekTime int = 60 * 60 * 24 * 7

func check(c *gin.Context) {
	userRaw, ok := c.Get("user")

	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user := userRaw.(*user.User)
	expiration := int64(c.MustGet("token_expire").(float64))

	t, err := token.Generate(user)

	if err != nil {
		log.Println(err)
	}

	c.SetCookie("token", t, weekTime, "/", "", false, true)
	c.JSON(http.StatusOK, common.JSON{
		"token": t,
	})
}

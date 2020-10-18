package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrtroian/notes/internal/api/internal/request"
	"github.com/mrtroian/notes/internal/hash"
	"github.com/mrtroian/notes/internal/token"
	"github.com/mrtroian/notes/internal/user"
)

func signin(c *gin.Context) {
	// db := c.MustGet("db").(*gorm.DB)
	body := request.NewLogin()
	defer body.Reset()

	if err := c.BindJSON(body); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := user.FindByUsername(body.Username)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := hash.Validate(body.Password, u.PasswordHash); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	t, err := token.Generate(u)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.SetCookie("token", t, weekTime, "/", "", false, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user":  u.Serialize(),
		"token": t,
	})
}

package auth

import (
	"log"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/mrtroian/notes/internal/hash"
	"github.com/mrtroian/notes/internal/request"
	"github.com/mrtroian/notes/internal/token"
)

func signup(c *gin.Context) {
	body := request.NewRegister()
	defer body.Reset()

	if err := c.BindJSON(body); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if user.ExistsWithUsername(body.Username) {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	hs, err := hash.Generate(body.Password)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	u := user.NewUser(body.Username, body.Email, hs)

	err := user.Save(u)

	if err != nil {
		log.Println(err)
		// Actually, any server side error
		c.AbortWithStatus(http.StatusTeapot)
		return
	}

	t, err := token.Generate(user)

	if err != nil {
		log.Println(err)
	}

	// Setting cookies is not RESTful, but who cares??
	c.SetCookie("token", t, weekTime, "/", "", false, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": t,
	})
}

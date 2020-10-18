package notes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrtroian/notes/internal/api/internal/request"
	"github.com/mrtroian/notes/internal/note"
)

func create(c *gin.Context) {
	body := request.NewCreateNote()
	defer body.Reset()

	if err := c.BindJSON(&body); err != nil {
		log.Println("api/v1/notes/create", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := readUser(c)

	if err != nil {
		log.Println("api/v1/notes/create", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	n := &note.Note{
		Title: body.Title,
		Text:  body.Text,
		User:  u,
	}

	err = n.Create()

	if err != nil {
		log.Println("api/v1/notes/create", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, n)
}

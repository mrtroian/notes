package notes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrtroian/notes/internal/note"
)

func readAll(c *gin.Context) {
	u, err := readUser(c)

	if err != nil {
		log.Println("api/v1/notes/create", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	notes, err := note.LoadByUserID(u.ID, 10, 0)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, notes)
}

func readOne(c *gin.Context) {
	u, err := readUser(c)

	if err != nil {
		log.Println("api/v1/notes/create", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println("api/v1/notes/create", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	n, err := note.Load(u.ID, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, n)
}

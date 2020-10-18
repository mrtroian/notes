package notes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrtroian/notes/internal/note"
)

func delete(c *gin.Context) {
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

	if err := note.Delete(u.ID, id); err != nil {
		if err == note.ErrForbidden {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

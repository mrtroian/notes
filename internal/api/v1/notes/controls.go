package notes

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/mrtroian/notes/internal/common"
    "github.com/mrtroian/notes/internal/database/models"
    "github.com/mrtroian/notes/internal/request"
)

// create a new note
func create(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    body := new(request.Create)

    if err := c.BindJSON(&body); err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    user := c.MustGet("user").(models.User)
    note := models.Note{
        Title: body.Title,
        Text:  body.Text,
        User:  user,
    }

    db.NewRecord(note)
    db.Create(&note)

    c.JSON(http.StatusOK, note.Serialize())
}

// get all notes
func retrieveAll(c *gin.Context) {
    var notes []models.Note

    db := c.MustGet("db").(*gorm.DB)
    cursor := c.Query("cursor")
    recent := c.Query("recent")

    if cursor == "" {
        if err := db.Preload("User").Limit(10).Order("id desc").Find(&notes).Error; err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }
    } else {
        cond := "id < ?"

        if recent == "1" {
            cond = "id > ?"
        }

        if err := db.Preload("User").Limit(10).Order("id desc").Where(cond, cursor).Find(&notes).Error; err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }
    }

    length := len(notes)
    serialized := make([]common.JSON, length, length)

    for i := 0; i < length; i++ {
        serialized[i] = notes[i].Serialize()
    }

    c.JSON(http.StatusOK, serialized)
}

// get note by id
func retrieveOne(c *gin.Context) {
    note := new(models.Note)
    db := c.MustGet("db").(*gorm.DB)
    id := c.Param("id")

    if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(note).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        return
    }

    c.JSON(http.StatusOK, note.Serialize())
}

// update note by id
func update(c *gin.Context) {
    // @TODO: implement update
    panic("Not implemented")
}

// delete note by id
func delete(c *gin.Context) {
    note := new(models.Note)
    user := c.MustGet("user").(models.User)
    db := c.MustGet("db").(*gorm.DB)
    id := c.Param("id")

    if err := db.Where("id = ?", id).First(note).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        return
    }

    if note.UserID != user.ID {
        c.AbortWithStatus(http.StatusForbidden)
        return
    }

    db.Delete(note)
    c.Status(http.StatusNoContent)
}

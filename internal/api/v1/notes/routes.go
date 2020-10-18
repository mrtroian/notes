package notes

import (
    "github.com/gin-gonic/gin"
    "github.com/mrtroian/notes/internal/api/internal/middleware"
)

func ApplyRoutes(r *gin.RouterGroup) {
    notes := r.Group("/notes", middleware.IsAuthorised)
    {
        notes.POST("/", create)
        notes.GET("/", readAll)
        notes.GET("/:id", readOne)
        notes.PATCH("/:id", update)
        notes.DELETE("/:id", delete)
    }
}

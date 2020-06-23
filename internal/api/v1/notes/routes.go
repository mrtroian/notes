package notes

import (
    "github.com/gin-gonic/gin"
    "github.com/mrtroian/notes/internal/middleware"
)

func ApplyRoutes(r *gin.RouterGroup) {
    notes := r.Group("/notes")
    {
        notes.POST("/", middleware.IsAuthorised, create)
        notes.GET("/", middleware.IsAuthorised, retrieveAll)
        notes.GET("/:id", middleware.IsAuthorised, retrieveOne)
        notes.PATCH("/:id", middleware.IsAuthorised, update)
        notes.DELETE("/:id", middleware.IsAuthorised, delete)
    }
}

package apiv1

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/mrtroian/notes/internal/api/v1/auth"
    "github.com/mrtroian/notes/internal/api/v1/notes"
)

func ping(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "success!",
    })
}

// Apply routes to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
    v1 := r.Group("/v1")
    {
        v1.GET("/ping", ping)
        auth.ApplyRoutes(v1)
        notes.ApplyRoutes(v1)
    }
}

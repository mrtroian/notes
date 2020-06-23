package api

import (
    "github.com/gin-gonic/gin"
    apiv1 "github.com/mrtroian/notes/internal/api/v1"
)

func ApplyRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        apiv1.ApplyRoutes(api)
    }
}

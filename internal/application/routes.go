package application

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Serve static files
func ApplyRoute(r *gin.Engine) {
	r.Use(static.Serve("/", static.LocalFile("./web/", true)))
}

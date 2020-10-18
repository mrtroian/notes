package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrtroian/notes/internal/api/internal/middleware"
	apiv1 "github.com/mrtroian/notes/internal/api/v1"
	"github.com/mrtroian/notes/internal/config"
)

var srv = &http.Server{
	Addr: fmt.Sprintf("%s:%s", config.GetHost(), config.GetPort()),
}

func Start(r *gin.Engine) {
	api := r.Group("/api")
	{
		// append here other routers
		// i.e. v2 or proto
		apiv1.ApplyRoutes(api)
	}

	r.Use(middleware.JWTMiddleware())
	srv.Handler = r

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server: %s\n", err)
		}
	}()
}

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)

	fmt.Print("\r")

	if err != nil {
		log.Println("server: forced shutdown:", err)
	} else {
		log.Println("server: shutdown correctly")
	}
}

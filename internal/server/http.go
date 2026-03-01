package server

import (
	"github.com/gin-gonic/gin"
	"github.com/khy20040121/orbit-layout-basic/internal/handler"
	"github.com/khy20040121/orbit-layout-basic/internal/middleware"
	"github.com/khy20040121/orbit-layout-basic/pkg/helper/resp"
	"github.com/khy20040121/orbit-layout-basic/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	userHandler *handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Orbit!",
		})
	})
	r.GET("/user", userHandler.GetUserById)

	return r
}

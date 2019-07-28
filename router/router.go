package router

import (
	"net/http"

	"github.com/hyper-xx/mygin/handler/monitor"
	"github.com/hyper-xx/mygin/router/middleware"

	"github.com/gin-gonic/gin"
)

//Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	//404handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	//the health check handlers
	svcd := g.Group("/monitor")
	{
		svcd.GET("/health", monitor.HealthCheck)
		svcd.GET("/disk", monitor.DiskCheck)
		svcd.GET("cpu", monitor.CPUCheck)
		svcd.GET("ram", monitor.RAMCheck)
	}

	return g
}

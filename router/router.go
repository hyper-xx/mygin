package router

import (
	"net/http"

	"github.com/hyper-xx/mygin/handler/monitor"
	"github.com/hyper-xx/mygin/handler/user"
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

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

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

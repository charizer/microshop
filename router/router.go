package router

import (
	"github.com/gin-gonic/gin"
	"microshop/application"
	"microshop/router/middleware"
	"net/http"
)

func Load(g *gin.Engine) *gin.Engine {
	// Middlewares
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	authGroup := g.Group("/api/v1/auth")
	authGroup.Use(middleware.StrictTokenHandler())
	{
		authGroup.GET("/cart/list", application.CartList)
	}
	noAuthGroup := g.Group("/api/v1/noauth")
	{
		noAuthGroup.GET("/list", application.CartList)
	}
	return g
}

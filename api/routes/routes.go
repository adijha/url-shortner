package routes

import (
	"github.com/adijha/url-shortner/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// static UI routes
	InitializeStaticRoutes(router)

	//url shortener routes
	router.GET("/:url", controllers.ResolveURL)
	router.POST("/api/shorten", controllers.ShortenURL)

	//User routes
	InitializeUserRoutes(router)
}

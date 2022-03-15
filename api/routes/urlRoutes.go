package routes

import (
	"github.com/adijha/url-shortner/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeUrlRoutes(router *gin.Engine) {
	router.GET("/:url", controllers.ResolveURL)
	router.POST("/api/shorten", controllers.ShortenURL)
}

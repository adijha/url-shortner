package routes

import (
	"github.com/adijha/url-shortner/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeUrlRoutes(router *gin.Engine) {
	router.GET("/:url", controllers.ResolveURL)
	url := router.Group("/url")
	{
		url.POST("/shorten", controllers.ShortenURL)
		url.GET("/all", controllers.GetAllUrls)
		url.GET("/:user_id", controllers.GetUrlsByUser)
	}
}

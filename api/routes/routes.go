package routes

import (
	"net/http"

	"github.com/adijha/url-shortner/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// static UI routes
	router.Static("/assets", "./public/assets")
	router.LoadHTMLGlob("./public/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//url shortener routes
	router.GET("/:url", ResolveURL)
	router.POST("/api/shorten", ShortenURL)

	//User routes
	user := router.Group("/user")
	{
		user.GET("/", controllers.GetAllUsers)
		user.GET("/:id", controllers.GetUser)
		user.POST("/", controllers.CreateUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
}

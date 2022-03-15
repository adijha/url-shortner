package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeStaticRoutes(router *gin.Engine) {
	router.Static("/assets", "./view/assets")
	router.LoadHTMLGlob("./view/index.html")
	router.StaticFile("/favicon.ico", "./view/assets/favicon.ico")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

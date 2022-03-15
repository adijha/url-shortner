package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeStaticRoutes(router *gin.Engine) {
	router.Static("/assets", "./view/assets")
	router.LoadHTMLGlob("./view/index.html")
	//TODO: add favicon
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

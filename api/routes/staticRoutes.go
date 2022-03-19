package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func InitializeStaticRoutes(router *gin.Engine) {
// 	router.Static("/assets", "./view/assets")
// 	router.LoadHTMLGlob("./view/index.html")
// 	router.StaticFile("/favicon.ico", "./view/assets/favicon.ico")
// 	//show homepage
// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", nil)
// 	})
// }

func InitializeStaticRoutes(router *gin.Engine) {
	router.Static("/assets", "./view/dist/assets")
	router.LoadHTMLGlob("./view/dist/index.html")
	router.StaticFile("/favicon.ico", "./view/dist/favicon.ico")
	//show homepage
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

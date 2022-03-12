package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adijha/url-shortner/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()
	router.Static("/assets", "./public/assets")
	router.LoadHTMLGlob("./public/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/:url", routes.ResolveURL)
	router.POST("/api/shorten", routes.ShortenURL)
	router.Run(os.Getenv("PORT"))
}

package main

import (
	"fmt"
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
	router.Static("/home", "./public")
	router.GET("/:url", routes.ResolveURL)
	router.POST("/api/shorten", routes.ShortenURL)
	router.Run(os.Getenv("PORT"))
}

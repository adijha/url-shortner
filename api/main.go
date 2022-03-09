package main

import (
	"github.com/adijha/url-shortner/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/home", "./public")
	router.GET("/:url", routes.ResolveURL)
	router.POST("/api/shorten", routes.ShortenURL)
	router.Run(":80")
}

// func setupRoutes(app *fiber.App) {

// 	app.Static("/", "./public")
// 	app.Get("/:url", routes.ResolveURL)
// 	app.Post("/api/shorten", routes.ShortenURL)
// }

// err := godotenv.Load()
// if err != nil {
// 	fmt.Println(err)
// }

// app := fiber.New()

// app.Use(logger.New())
// app.Use(cors.New())

// setupRoutes(app)

// log.Fatal(app.Listen(os.Getenv("APP_PORT")))

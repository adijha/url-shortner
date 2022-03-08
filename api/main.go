package main

import (
	"github.com/adijha/url-shortner/routes"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Static("/", "./public")
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/shorten", routes.ShortenURL)
}

func main() {

	router := gin.Default()
	router.Static("/", "./public")
	// router.GET("/add/:x/:y", add)
	router.Run(":80")

	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// app := fiber.New()

	// app.Use(logger.New())
	// app.Use(cors.New())

	// setupRoutes(app)

	// log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

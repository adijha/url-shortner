package routes

import (
	"github.com/adijha/url-shortner/database"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func ResolveURL(c *gin.Context) {
	// get the short URL from the request
	shortURL := c.Param("url")

	// get the redis client
	r := database.CreateClient(0)
	defer r.Close()

	// get the URL from the redis database
	val, err := r.Get(database.Ctx, shortURL).Result()
	if err == redis.Nil {
		c.JSON(404, gin.H{
			"error": "URL not found!",
		})
		return
	} else if err != nil {
		c.JSON(500, gin.H{
			"error": "Cannot connect to database!",
		})
		return
	}

	c.Redirect(301, val)
}

// func ResolveURL(c *fiber.Ctx) error {
// 	url := c.Params("url")

// 	r := database.CreateClient(0)
// 	defer r.Close()

// 	value, err := r.Get(database.Ctx, url).Result()

// 	if err == redis.Nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Short URL not found in the database!"})
// 	} else if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot connect to database!"})
// 	}

// 	rInr := database.CreateClient(1)

// 	defer rInr.Close()

// 	_ = rInr.Incr(database.Ctx, "counter")

// 	return c.Redirect(value, 301)
// }

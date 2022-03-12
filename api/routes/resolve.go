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

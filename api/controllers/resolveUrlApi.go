package controllers

import (
	"github.com/adijha/url-shortner/cache"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func ResolveURL(c *gin.Context) {
	// get the short URL from the request
	shortURL := c.Param("url")

	// get the redis client
	r := cache.CreateClient(0)
	defer r.Close()

	// get the URL from the redis cache
	val, err := r.Get(cache.Ctx, shortURL).Result()
	if err == redis.Nil {
		//TODO: redirect to homepage after a delay, and show delay counter
		c.JSON(404, gin.H{
			"error": "URL not found!",
		})
		return
	} else if err != nil {
		c.JSON(500, gin.H{
			"error": "Cannot connect to cache!",
		})
		return
	}

	c.Redirect(301, val)
}

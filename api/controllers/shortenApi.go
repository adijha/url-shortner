package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adijha/url-shortner/cache"
	"github.com/adijha/url-shortner/database"
	"github.com/adijha/url-shortner/helpers"
	"github.com/adijha/url-shortner/request"
	"github.com/adijha/url-shortner/response"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func ShortenURL(c *gin.Context) {
	var req request.Url
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body!",
		})
		return
	}

	// implementing rate limiting
	r2 := cache.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(cache.Ctx, c.ClientIP()).Result()
	fmt.Println(val, "val")
	if err == redis.Nil {
		_ = r2.Set(cache.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		valInt, _ := strconv.Atoi(val)
		fmt.Println(valInt, "valInt")
		if valInt <= 0 {
			limit, _ := r2.TTL(cache.Ctx, c.ClientIP()).Result()
			c.JSON(400, gin.H{
				"error":            "Rate limit exceeded!",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
			return
		}
	}

	//verify url
	if !govalidator.IsURL(req.URL) {
		c.JSON(400, gin.H{
			"error": "Invalid URL!",
		})
		return
	}

	// 	check for domain error
	if !helpers.RemoveDomainError(req.URL) {
		c.JSON(500, gin.H{
			"error": "You cannot shorten this domain!",
		})
		return
	}

	req.URL = helpers.EnforceHTTP(req.URL)

	var id string

	if req.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = req.CustomShort
	}

	r := cache.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(cache.Ctx, id).Result()

	if val != "" {
		c.JSON(403, gin.H{
			"error": "Your custom short URL is already in use!",
		})
		return
	}

	if req.Expiry == 0 {
		req.Expiry = 24
	}

	err = r.Set(cache.Ctx, id, req.URL, req.Expiry*3600*time.Second).Err()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	// update on postgres db also
	database.DB.Create(&req)

	// create custom response
	resp := response.Url{
		URL:             req.URL,
		CustomShort:     "",
		Expiry:          req.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
		User_id:         req.User_id,
	}
	r2.Decr(cache.Ctx, c.ClientIP())

	val, _ = r2.Get(cache.Ctx, c.ClientIP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := r2.TTL(cache.Ctx, c.ClientIP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	c.JSON(200, resp)
}

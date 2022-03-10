package routes

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adijha/url-shortner/database"
	"github.com/adijha/url-shortner/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *gin.Context) {
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body!",
		})
		return
	}

	// 	// implementing rate limiting
	r2 := database.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
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

	req.URL = helpers.EnforceHTTP(req.URL)

	var id string

	if req.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = req.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()

	if val != "" {
		c.JSON(403, gin.H{
			"error": "Your custom short URL is already in use!",
		})
		return
	}

	if req.Expiry == 0 {
		req.Expiry = 24
	}

	fmt.Printf("%+v\n", req)

	if r.Exists(database.Ctx, req.CustomShort).Val() != 0 {
		c.JSON(400, gin.H{
			"error": "Custom short URL already exists!",
		})
		return
	}

	// 	// check for domain error
	if !helpers.RemoveDomainError(req.URL) {
		c.JSON(500, gin.H{
			"error": "You cannot shorten this domain!",
		})
		return
	}

	err = r.Set(database.Ctx, id, req.URL, req.Expiry*3600*time.Second).Err()

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Cannot connect to server",
		})
		return
	}

	resp := response{
		URL:             req.URL,
		CustomShort:     "",
		Expiry:          req.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
	}
	r2.Decr(database.Ctx, c.ClientIP())

	val, _ = r2.Get(database.Ctx, c.ClientIP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	// if r.SetNX(database.Ctx, req.CustomShort, req.URL, req.Expiry).Val() {
	// 	c.JSON(201, response{
	// 		URL:             req.URL,
	// 		CustomShort:     req.CustomShort,
	// 		Expiry:          req.Expiry,
	// 		XRateRemaining:  10,
	// 		XRateLimitReset: 10,
	// 	})
	// 	return
	// }

	c.JSON(200, resp)
}

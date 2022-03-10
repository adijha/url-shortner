package routes

import (
	"fmt"
	"time"

	"github.com/adijha/url-shortner/database"
	"github.com/adijha/url-shortner/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
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

	//verify url

	if !govalidator.IsURL(req.URL) {
		c.JSON(400, gin.H{
			"error": "Invalid URL!",
		})
		return
	}

	r2 := database.CreateClient(0)
	defer r2.Close()

	if req.Expiry == 0 {
		req.Expiry = time.Hour * 24
	}

	if req.CustomShort == "" {
		req.CustomShort = uuid.New().String()
	}

	fmt.Printf("%+v\n", req)

	if r2.Exists(database.Ctx, req.CustomShort).Val() != 0 {
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

	if r2.SetNX(database.Ctx, req.CustomShort, req.URL, req.Expiry).Val() {
		c.JSON(201, response{
			URL:             req.URL,
			CustomShort:     req.CustomShort,
			Expiry:          req.Expiry,
			XRateRemaining:  10,
			XRateLimitReset: 10,
		})
		return
	}

	c.JSON(200, gin.H{
		"Success": "Done everything!",
	})
}

// func ShortenURL(c *fiber.Ctx) error {
// 	body := new(request)

// 	if err := c.BodyParser(&body); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
// 	}

// 	// implementing rate limiting
// 	r2 := database.CreateClient(1)
// 	defer r2.Close()

// 	val, err := r2.Get(database.Ctx, c.IP()).Result()
// 	if err == redis.Nil {
// 		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
// 	} else {
// 		valInt, _ := strconv.Atoi(val)
// 		if valInt <= 0 {
// 			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
// 			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
// 				"error":            "Rate limit exceeded!",
// 				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
// 			})
// 		}
// 	}

// 	// check if the input is an actual URL
// 	if !govalidator.IsURL(body.URL) {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL!"})
// 	}

// 	// check for domain error
// 	if !helpers.RemoveDomainError(body.URL) {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You can't access it! :)"})
// 	}

// 	// enforce https, SSL
// 	body.URL = helpers.EnforceHTTP(body.URL)

// 	// custom URL by user functionality
// 	var id string

// 	if body.CustomShort == "" {
// 		id = uuid.New().String()[:6]
// 	} else {
// 		id = body.CustomShort
// 	}

// 	r := database.CreateClient(0)
// 	defer r.Close()

// 	val, _ = r.Get(database.Ctx, id).Result()

// 	if val != "" {
// 		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Your custom short URL is already in use!"})
// 	}

// 	if body.Expiry == 0 {
// 		body.Expiry = 24
// 	}

// 	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to connect to server!"})
// 	}

// 	resp := response{
// 		URL:             body.URL,
// 		CustomShort:     "",
// 		Expiry:          body.Expiry,
// 		XRateRemaining:  10,
// 		XRateLimitReset: 30,
// 	}

// 	r2.Decr(database.Ctx, c.IP())

// 	val, _ = r2.Get(database.Ctx, c.IP()).Result()
// 	resp.XRateRemaining, _ = strconv.Atoi(val)

// 	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
// 	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

// 	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

// 	return c.Status(fiber.StatusOK).JSON(resp)
// }

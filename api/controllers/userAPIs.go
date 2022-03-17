package controllers

import (
	"fmt"

	"github.com/adijha/url-shortner/cache"
	"github.com/adijha/url-shortner/database"
	"github.com/adijha/url-shortner/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	d := database.DB.Where("id = ?", id).Delete(&user)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&user)
	database.DB.Save(&user)
	c.JSON(200, user)
}
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	database.DB.Create(&user)
	c.JSON(200, user)
}
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}
func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

type URls struct {
	User_id string `json:"user_id"`
}

func GetUrls(c *gin.Context) {

	var req URls
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body!",
		})
		return
	}
	//get user id
	// userID := c.Param("id")

	// get the redis client
	r := cache.CreateClient(0)
	defer r.Close()

	val, err := r.Get(cache.Ctx, req.User_id).Result()
	// val, err := r.Get(cache.Ctx, userID).Result()
	if err == redis.Nil {
		c.JSON(404, gin.H{
			"error": "No URL found for user",
			"id":    req.User_id,
		})
		return
	} else if err != nil {
		c.JSON(500, gin.H{
			"error": "Cannot connect to cache!",
		})
		return
	}

	c.JSON(200, val)
}

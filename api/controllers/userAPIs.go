package controllers

import (
	"fmt"

	"github.com/adijha/url-shortner/database"
	"github.com/adijha/url-shortner/models"
	"github.com/gin-gonic/gin"
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
func GetAllUrls(c *gin.Context) {
	var urls []models.Url
	if err := database.DB.Find(&urls).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, urls)
	}
}

func GetUrlsByUser(c *gin.Context) {
	// get the short User Id from the request
	userId := c.Param("user_id")
	// get the url from the database
	var urls []models.Url
	if err := database.DB.Where("user_id = ?", userId).Find(&urls).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, urls)
	}

	//get ot from redis cache
	// var keys []string
	// keys = append(keys, "foo")
	// keys = append(keys, "bar")
	// get the redis client
	// r := cache.CreateClient(0)
	// defer r.Close()

	// val, err := r.MGet(cache.Ctx, keys...).Result()
	// get urls from the redis cache
	// val, err := r.Get(cache.Ctx, userId).Result()
	// if err == redis.Nil {
	// 	//TODO: redirect to homepage after a delay, and show delay counter
	// 	c.JSON(404, gin.H{
	// 		"error": "URL not found!",
	// 	})
	// 	return
	// } else if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"error": "Cannot connect to cache!",
	// 	})
	// 	return
	// }

	// c.JSON(200, val)

}

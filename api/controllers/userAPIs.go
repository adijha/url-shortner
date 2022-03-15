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

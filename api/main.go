package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adijha/url-shortner/database"
	"github.com/adijha/url-shortner/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

func main() {
	// Load environment variables
	err = godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	// init postgres database
	db = database.InitDB()
	db.AutoMigrate(&User{})
	//router init
	router := gin.Default()
	// static UI routes
	router.Static("/assets", "./public/assets")
	router.LoadHTMLGlob("./public/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//url shortener routes
	router.GET("/:url", routes.ResolveURL)
	router.POST("/api/shorten", routes.ShortenURL)

	//User routes
	user := router.Group("/user")
	{
		user.GET("/", GetAllUsers)
		user.GET("/:id", GetUser)
		user.POST("/", CreateUser)
		user.PUT("/:id", UpdateUser)
		user.DELETE("/:id", DeleteUser)
	}

	// listen and serve
	router.Run(os.Getenv("PORT"))

	//close db
	dbSQL, err := db.DB()
	if err != nil {
		defer dbSQL.Close()
	}
	defer dbSQL.Close()
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	d := db.Where("id = ?", id).Delete(&user)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
func UpdateUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(200, user)
}
func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	db.Create(&user)
	c.JSON(200, user)
}
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}
func GetAllUsers(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

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

type Person struct {
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
	db.AutoMigrate(&Person{})
	//router init
	router := gin.Default()
	//url shortener routes
	router.Static("/assets", "./public/assets")
	router.LoadHTMLGlob("./public/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/:url", routes.ResolveURL)
	router.POST("/api/shorten", routes.ShortenURL)
	//postgres routes
	router.GET("/people/", GetPeople)
	router.GET("/people/:id", GetPerson)
	router.POST("/people", CreatePerson)
	router.PUT("/people/:id", UpdatePerson)
	router.DELETE("/people/:id", DeletePerson)
	router.Run(os.Getenv("PORT"))

	//close db
	dbSQL, err := db.DB()
	if err != nil {
		defer dbSQL.Close()
	}
	defer dbSQL.Close()
}

func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
func UpdatePerson(c *gin.Context) {
	var person Person
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	db.Save(&person)
	c.JSON(200, person)
}
func CreatePerson(c *gin.Context) {
	var person Person
	c.BindJSON(&person)
	db.Create(&person)
	c.JSON(200, person)
}
func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}
func GetPeople(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}
}

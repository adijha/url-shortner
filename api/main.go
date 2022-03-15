package main

import (
	"fmt"
	"os"

	"github.com/adijha/url-shortner/database"
	"github.com/adijha/url-shortner/models"
	"github.com/adijha/url-shortner/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// Load environment variables
	err = godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	// init postgres database
	db = database.InitDB()
	db.AutoMigrate(&models.User{})

	//router init
	router := gin.Default()
	routes.InitializeRoutes(router,&db)

	// listen and serve
	router.Run(os.Getenv("PORT"))

	//close db
	dbSQL, err := db.DB()
	if err != nil {
		defer dbSQL.Close()
	}
	defer dbSQL.Close()
}

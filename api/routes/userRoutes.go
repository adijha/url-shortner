package routes

import (
	"github.com/adijha/url-shortner/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/", controllers.GetAllUsers)
		user.GET("/:id", controllers.GetUser)
		user.POST("/", controllers.CreateUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
}

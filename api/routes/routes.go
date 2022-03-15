package routes

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// static UI routes
	InitializeStaticRoutes(router)
	//url shortener routes
	InitializeUrlRoutes(router)
	//User routes
	InitializeUserRoutes(router)
}

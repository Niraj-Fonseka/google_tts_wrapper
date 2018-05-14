package routers

import (
	"webapp_go/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRouters() *gin.Engine {
	Router = gin.Default()
	v1 := Router.Group("/")
	{
		v1.GET("/hello", controllers.GetHealth)
		v1.GET("/test", controllers.CreateCreds)
	}
	return Router
}

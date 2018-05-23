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
		v1.GET("/generate_report", controllers.GETDailyReport)
		v1.GET("/news/:source", controllers.GetBreakingNews)
		v1.POST("/decode", controllers.RunDailyReport)

	}
	return Router
}

package routers

import (
	"google_tts_wrapper/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRouters() *gin.Engine {
	Router = gin.Default()
	v1 := Router.Group("/")
	{
		v1.GET("/health", controllers.GetHealth)
		v1.GET("/generate_report", controllers.GETDailyReport)
		v1.GET("/news/:source", controllers.GetBreakingNews)

	}
	return Router
}

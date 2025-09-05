package app

import(
	"github.com/gin-gonic/gin"
	"github.com/Ruan6401/health/api/controllers"
)


func SetupMappings(engine *gin.Engine){
	engine.GET("/Marco",controllers.NewHealthController().Handle)
}
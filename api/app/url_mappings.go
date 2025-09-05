package app

import(
	"github.com/gin-gonic/gin"
	"github.com/Ruan6401/health/api/controllers"
	"github.com/Ruan6401/health/api/services"
)


func SetupMappings(engine *gin.Engine, hSvc services.HealthService){
	engine.GET("/Marco",controllers.NewHealthController(hSvc).Handle)
}
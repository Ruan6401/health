package app

import (
	"github.com/gin-gonic/gin"

	"github.com/Ruan6401/health/config"
)

type Application struct{
	Engine *gin.Engine
}

func NewApplication() *Application{
	engine := gin.Default()
	return &Application{
		Engine: engine,
	}
}
package app

import (
	"github.com/gin-gonic/gin"

	"github.com/Ruan6401/health/api/config"
)

type Application struct{
	Engine *gin.Engine
	Config *config.Config
}

func NewApplication(cfg *config.Config) *Application{
	engine := gin.Default()
	return &Application{
		Engine: engine,
		Config: cfg,
	}
}
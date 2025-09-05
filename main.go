package main

import (
	"fmt"
	"log"

	"github.com/Ruan6401/health/api/app"
	"github.com/Ruan6401/health/api/config"
)


func main(){
	cfg,err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config, %s",err)
	}
	application := app.NewApplication(cfg)
	app.SetupMappings(application.Engine)
	addr := fmt.Sprintf(":%d", cfg.Port)
	if err := application.Engine.Run(addr); err != nil {
		log.Fatalf("server run error: %v", err)
	}
}
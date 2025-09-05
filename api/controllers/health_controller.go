package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Ruan6401/health/api/services"
)

type healthController struct{
	svc services.HealthService
}

func NewHealthController(svc services.HealthService) *healthController{
	return &healthController{
		svc: svc,
	}
}

func (h *healthController)Handle(c *gin.Context){
	
	response := h.svc.AnswerMarco()

	c.JSON(http.StatusOK,response)
}
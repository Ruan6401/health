package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Ruan6401/health/api/domain"
)

type healthController struct{}

func NewHealthController() *healthController{
	return &healthController{}
}

func (h *healthController)Handle(c *gin.Context){
	response := &domain.Response{
		Resp: "Polo!",
	}

	c.JSON(http.StatusOK,response)
}
package services

import "github.com/Ruan6401/health/api/domain"

type HealthService interface{
	AnswerMarco() *domain.Response
}

type healthService struct{}

func NewHealthService() HealthService{
	return &healthService{}
}

func (h *healthService)AnswerMarco() *domain.Response{
	response := &domain.Response{
		Resp: "Polo!",
	}

	return response
}
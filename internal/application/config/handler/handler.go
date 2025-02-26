package handler

import "github.com/radityacandra/ecommerce-connector-cli/internal/application/config/service"

type Handler struct {
	Service service.IService
}

func NewHandler() *Handler {
	service := service.NewService()

	return &Handler{
		Service: service,
	}
}

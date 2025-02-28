package handler

import (
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/repository"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/service"
	"github.com/radityacandra/ecommerce-connector-cli/internal/core"
)

type Handler struct {
	Service service.IService
}

func NewHandler(deps *core.Dependency) *Handler {
	repository := repository.NewRepository(deps.Db)
	service := service.NewService(repository)

	return &Handler{
		Service: service,
	}
}

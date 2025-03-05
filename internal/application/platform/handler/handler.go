package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/repository"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/service"
	"github.com/radityacandra/ecommerce-connector-cli/internal/core"
)

type Handler struct {
	Service   service.IService
	Validator *validator.Validate
}

func NewHandler(deps *core.Dependency) *Handler {
	repository := repository.NewRepository(deps.Db)
	service := service.NewService(repository)

	return &Handler{
		Service:   service,
		Validator: deps.Validator,
	}
}

package handler

import (
	"github.com/go-playground/validator/v10"
	configService "github.com/radityacandra/ecommerce-connector-cli/internal/application/config/service"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/repository"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/service"
	"github.com/radityacandra/ecommerce-connector-cli/internal/core"
)

type Handler struct {
	UserService   service.IService
	ConfigService configService.IService
	Validator     *validator.Validate
}

func NewHandler(dep *core.Dependency) *Handler {
	repository := repository.NewRepository(dep.Db)
	service := service.NewService(repository)
	configService := configService.NewService()

	return &Handler{
		UserService:   service,
		ConfigService: configService,
		Validator:     dep.Validator,
	}
}

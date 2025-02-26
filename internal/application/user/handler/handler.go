package handler

import (
	configService "github.com/radityacandra/ecommerce-connector-cli/internal/application/config/service"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/repository"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/service"
	"github.com/radityacandra/ecommerce-connector-cli/internal/core"
)

type Handler struct {
	UserService   service.IService
	ConfigService configService.IService
}

func NewHandler(dep *core.Dependency) *Handler {
	repository := repository.NewRepository(dep.Db)
	service := service.NewService(repository)
	configService := configService.NewService()

	return &Handler{
		UserService:   service,
		ConfigService: configService,
	}
}

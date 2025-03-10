package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/client/fakestore"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/repository"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/service"
	"github.com/radityacandra/ecommerce-connector-cli/internal/core"
)

type Handler struct {
	Service   service.IService
	Validator *validator.Validate
}

func NewHandler(deps *core.Dependency) *Handler {
	repository := repository.NewRepository(deps.Db)
	fakestoreClient := fakestore.NewFakestore("https://fakestoreapi.com")
	service := service.NewService(repository, fakestoreClient)

	return &Handler{
		Service:   service,
		Validator: deps.Validator,
	}
}

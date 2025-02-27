package service

import (
	"context"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/repository"
)

type IService interface {
	Identify(ctx context.Context, email string) (IdentifyOutput, error)
	Register(ctx context.Context, input RegisterInput) error
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}

package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/model"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/repository"
)

type SetPlatformInput struct {
	UserId       string `validate:"required,uuid"`
	PlatformType string `validate:"required,oneof=native shopee tokopedia"`
}

var (
	ErrConflictingPlatform = errors.New("platform already exist")
)

func (s *Service) Set(ctx context.Context, input SetPlatformInput) error {
	// find by userId and platformType
	_, err := s.Repository.FindPlatformByUserIdAndType(ctx, repository.FindByUserIdAndTypeInput{
		UserId: input.UserId,
		Type:   input.PlatformType,
	})
	if err == nil {
		err = ErrConflictingPlatform
		return err
	}

	platform := model.NewPlatform(uuid.NewString(), input.UserId, input.PlatformType)

	err = s.Repository.Insert(ctx, platform)

	return err
}

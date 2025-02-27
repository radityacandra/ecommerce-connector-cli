package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/model"
)

type RegisterInput struct {
	Email          string `validate:"required,email"`
	Name           string `validate:"required"`
	ProfilePicture string
}

var (
	ErrConflictingUser = errors.New("user with the same email address has already exist")
)

func (s *Service) Register(ctx context.Context, input RegisterInput) error {
	// find user by email
	_, err := s.Repository.FindUserByEmail(ctx, input.Email)
	if err == nil {
		err = errors.Join(err, ErrConflictingUser)
		return err
	}

	user := model.NewUser(uuid.NewString(), input.Name, input.Email, input.ProfilePicture)

	err = s.Repository.Insert(ctx, user)

	return err
}

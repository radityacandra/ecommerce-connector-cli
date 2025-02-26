package service

import (
	"context"
)

type IdentifyOutput struct {
	Email    string
	UserName string
	UserId   string
}

func (s *Service) Identify(ctx context.Context, email string) (output IdentifyOutput, err error) {
	// find user by email
	user, err := s.Repository.FindUserByEmail(ctx, email)
	if err != nil {
		return
	}

	output.Email = user.Email
	output.UserId = user.Id
	output.UserName = user.Name

	return
}

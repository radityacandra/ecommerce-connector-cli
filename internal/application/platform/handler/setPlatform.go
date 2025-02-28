package handler

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/service"
)

func (h *Handler) SetPlatform(ctx context.Context, userId, platformType string) error {
	input := service.SetPlatformInput{
		UserId:       userId,
		PlatformType: platformType,
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(input)
	if err != nil {
		return err
	}

	err = h.Service.Set(ctx, input)
	if err == nil {
		fmt.Println("platform set successfully")
	}

	return err
}

package handler

import (
	"context"
	"fmt"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/service"
)

func (h *Handler) SetPlatform(ctx context.Context, userId, platformType string) error {
	input := service.SetPlatformInput{
		UserId:       userId,
		PlatformType: platformType,
	}

	err := h.Validator.Struct(input)
	if err != nil {
		return err
	}

	err = h.Service.Set(ctx, input)
	if err == nil {
		fmt.Println("platform set successfully")
	}

	return err
}

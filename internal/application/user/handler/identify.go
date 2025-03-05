package handler

import (
	"context"
	"fmt"
)

func (h *Handler) Identify(ctx context.Context) error {
	fmt.Print("please provide your email address: ")

	var email string
	fmt.Scan(&email)

	if err := h.Validator.Var(email, "required,email"); err != nil {
		return err
	}

	user, err := h.UserService.Identify(ctx, email)
	if err != nil {
		return err
	}

	err = h.ConfigService.InitConfig(ctx, "user_id", user.UserId)
	if err == nil {
		fmt.Println("provile saved successfully")
	}

	return err
}

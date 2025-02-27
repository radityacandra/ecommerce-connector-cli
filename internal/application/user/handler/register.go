package handler

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/service"
)

func (h *Handler) Register(ctx context.Context) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	var input service.RegisterInput
	firstValidate := true
	for {
		err = validate.StructCtx(ctx, input)
		if err == nil {
			break
		}
		if !firstValidate {
			fmt.Println("validation error: ", err)
		}

		fmt.Print("Please input email address: ")
		fmt.Scan(&input.Email)

		fmt.Print("Please input your name: ")
		fmt.Scanln(&input.Name)

		firstValidate = false
	}

	err = h.UserService.Register(ctx, input)
	if err == nil {
		fmt.Println("success registering a new user!")
	}

	return
}

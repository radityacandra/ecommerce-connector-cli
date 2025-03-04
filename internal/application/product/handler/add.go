package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
	"github.com/radityacandra/ecommerce-connector-cli/pkg/io"
)

func (h *Handler) Add(ctx context.Context, userId string) error {
	var product = model.Product{
		UserId: userId,
	}

	fmt.Print("provide product name: ")
	product.Name = io.ScanString(os.Stdin)

	fmt.Print("provide product description: ")
	product.Description = io.ScanString(os.Stdin)

	fmt.Print("provide product price: ")
	product.Price = io.ScanInt(os.Stdin)

	fmt.Print("provide product category: ")
	product.Category = io.ScanString(os.Stdin)

	fmt.Print("provide product ean code: ")
	product.EanCode = io.ScanString(os.Stdin)

	err := h.Service.Add(ctx, product)
	if err == nil {
		fmt.Println("product has been added")
	}

	return err
}

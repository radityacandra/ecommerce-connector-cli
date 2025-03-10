package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
	"github.com/radityacandra/ecommerce-connector-cli/pkg/io"
)

func (h *Handler) Add(ctx context.Context, product model.Product) error {
	if product.Name == "" {
		fmt.Print("provide product name: ")
		product.Name = io.ScanString(os.Stdin)
	}

	if product.Description == "" {
		fmt.Print("provide product description: ")
		product.Description = io.ScanString(os.Stdin)
	}

	if product.Price == 0 {
		fmt.Print("provide product price: ")
		product.Price = io.ScanInt(os.Stdin)
	}

	if product.Category == "" {
		fmt.Print("provide product category: ")
		product.Category = io.ScanString(os.Stdin)
	}

	if product.EanCode == "" {
		fmt.Print("provide product ean code: ")
		product.EanCode = io.ScanString(os.Stdin)
	}

	if err := h.Validator.Struct(product); err != nil {
		return err
	}

	err := h.Service.Add(ctx, product)
	if err == nil {
		fmt.Println("product has been added")
	}

	return err
}

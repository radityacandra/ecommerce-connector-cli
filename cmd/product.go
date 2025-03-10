/*
Copyright © 2025 Raditya Chandra <radityacandra@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/handler"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
	"github.com/radityacandra/ecommerce-connector-cli/internal/core"
	"github.com/radityacandra/ecommerce-connector-cli/pkg/database"
	"github.com/spf13/cobra"
)

// productCmd represents the product command
var productCmd = &cobra.Command{
	Use:   "product",
	Short: "miscellaneous operations of product",
	Long: `Don't use this command allone. use with subcommand:
	
	- add: add new product`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please run product subcommand, see available subcommand with --help")
	},
}

var addProductCmd = &cobra.Command{
	Use:   "add",
	Short: "Add products and perform integration to registered platform",
	Long: `This feature allows you to manage product details, pricing, and availability 
	while ensuring smooth synchronization across multiple sales channels such as Shopee, Tokopedia, and more.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		c := handleConfig(ctx)
		db, err := database.NewDatabase(ctx, c.FirebaseCredentialsFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}

		if c.UserId == "" {
			fmt.Println("userId not found. please run identify command first, or you can register with running register command")
			os.Exit(1)
			return
		}

		dep := core.NewDependency(db, validator.New(validator.WithRequiredStructEnabled()))

		handler := handler.NewHandler(dep)
		product.UserId = c.UserId
		err = handler.Add(ctx, product)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var product model.Product

func init() {
	productCmd.AddCommand(addProductCmd)
	rootCmd.AddCommand(productCmd)

	addProductCmd.Flags().StringVarP(&product.Name, "product-name", "n", "", "specify product name")
	addProductCmd.Flags().StringVarP(&product.Description, "product-description", "d", "", "specify product description")
	addProductCmd.Flags().IntVarP(&product.Price, "product-price", "p", 0, "specify product price")
	addProductCmd.Flags().StringVarP(&product.Category, "product-category", "c", "", "specify product category")
	addProductCmd.Flags().StringVarP(&product.EanCode, "product-ean", "e", "", "specify product ean code")
}

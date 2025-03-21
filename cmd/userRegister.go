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
	configHandler "github.com/radityacandra/ecommerce-connector-cli/internal/application/config/handler"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/handler"
	"github.com/radityacandra/ecommerce-connector-cli/internal/core"
	"github.com/radityacandra/ecommerce-connector-cli/pkg/config"
	"github.com/radityacandra/ecommerce-connector-cli/pkg/database"
	"github.com/spf13/cobra"
)

// userRegisterCmd represents the userRegister command
var userRegisterCmd = &cobra.Command{
	Use:   "user-register",
	Short: "Register new user to system",
	Long:  `Register new user to the system. you can use this user email on any machine later`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		var c *core.Config
		var err error
		for i := 0; i < 2; i++ {
			c, err = config.LoadConfig()
			if err == nil {
				break
			}

			fmt.Println("required config not found, initiating config")
			configHandler := configHandler.NewHandler()
			err = configHandler.InitConfig(ctx, "firebase_credentials_file_path")
			if err != nil {
				fmt.Println("error occured: ", err)
			}
		}
		if err != nil || c == nil {
			fmt.Println("failed to initiate config with error: ", err)
			os.Exit(1)
			return
		}

		db, err := database.NewDatabase(ctx, c.FirebaseCredentialsFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}

		dep := core.NewDependency(db, validator.New(validator.WithRequiredStructEnabled()))

		handler := handler.NewHandler(dep)
		err = handler.Register(ctx)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(userRegisterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userRegisterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userRegisterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

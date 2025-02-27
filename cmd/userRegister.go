/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

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

		dep := core.NewDependency(db)

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

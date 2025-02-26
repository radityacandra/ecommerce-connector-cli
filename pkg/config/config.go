package config

import (
	"context"
	"errors"
	"os"

	"github.com/radityacandra/ecommerce-connector-cli/internal/core"
	"github.com/spf13/viper"
)

func LoadConfig() (*core.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.ecommerce-connector-cli")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c core.Config
	viper.Unmarshal(&c)

	return &c, nil
}

func SaveConfig(ctx context.Context, config *core.Config) error {
	home, _ := os.UserHomeDir()
	configPath := home + "/.ecommerce-connector-cli/config.yaml"

	// check config file existence, create if not exist
	_, err := os.Stat(configPath)
	if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(configPath)
		if err != nil {
			return err
		}

		file.Close()
	} else if err != nil {
		return err
	}

	for key, val := range config.ToMap() {
		viper.Set(key, val)
	}

	viperConfiguration()

	err = viper.WriteConfig()

	return err
}

func viperConfiguration() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.ecommerce-connector-cli")
}

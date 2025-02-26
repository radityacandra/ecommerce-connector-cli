package service

import (
	"context"

	"github.com/radityacandra/ecommerce-connector-cli/internal/core"
	"github.com/radityacandra/ecommerce-connector-cli/pkg/config"
)

func (s *Service) InitConfig(ctx context.Context, configName string, configValue interface{}) error {
	// check if config exists
	c, err := config.LoadConfig()
	if err != nil {
		c = &core.Config{}
	}

	err = c.SetConfigField(configName, configValue)
	if err != nil {
		return err
	}

	err = config.SaveConfig(ctx, c)

	return err
}

package service

import "context"

type IService interface {
	InitConfig(ctx context.Context, configName string, configValue interface{}) (err error)
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

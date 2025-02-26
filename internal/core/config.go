package core

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrCanSetValue   = errors.New("failed to set value")
	ErrInvalidType   = errors.New("missmatched value type")
	ErrFieldNotFound = errors.New("field not found")
)

type Config struct {
	FirebaseCredentialsFilePath string `mapstructure:"firebase_credentials_file_path"`
	UserId                      string `mapstructure:"user_id"`
}

func (c *Config) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(c).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if tagValue, ok := field.Tag.Lookup("mapstructure"); ok {
			result[tagValue] = val.Field(i).Interface()
		}
	}

	return result
}

func (c *Config) SetConfigField(configName string, configValue any) error {
	val := reflect.ValueOf(c).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if tagValue, ok := field.Tag.Lookup("mapstructure"); ok && tagValue == configName {
			fieldValue := val.Field(i)
			if !fieldValue.CanSet() {
				ErrCanSetValue = fmt.Errorf("cannot set field %s", configName)
				return ErrCanSetValue
			}

			newValue := reflect.ValueOf(configValue)
			if !newValue.Type().AssignableTo(fieldValue.Type()) {
				ErrInvalidType = fmt.Errorf("invalid type for field %s: expected %v, got %v",
					configName, fieldValue.Type(), newValue.Type())
				return ErrInvalidType
			}

			fieldValue.Set(newValue)
			return nil
		}
	}

	ErrFieldNotFound = fmt.Errorf("field %s not found", configName)
	return ErrFieldNotFound
}

package handler

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type InitConfigInput struct {
}

func (h *Handler) InitConfig(ctx context.Context, configName string) error {
	fmt.Printf("please input value for %s: ", configName)

	reader := bufio.NewReader(os.Stdin)
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		return err
	}
	bytes = bytes[:len(bytes)-1]

	jsonByte, _ := json.Marshal(string(bytes))

	var configValue interface{}
	if err := json.Unmarshal(jsonByte, &configValue); err != nil {
		return err
	}

	err = h.Service.InitConfig(ctx, configName, configValue)
	if err == nil {
		fmt.Println("config saved successfully")
	}

	return err
}

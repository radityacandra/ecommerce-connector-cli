package handler_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/handler"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/service"
	mockService "github.com/radityacandra/ecommerce-connector-cli/mocks/internal_/application/product/service"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Add(t *testing.T) {
	type fields struct {
		Service service.IService
	}
	type args struct {
		ctx     context.Context
		product model.Product
	}
	type test struct {
		name    string
		fields  fields
		mock    func(test) test
		args    args
		wantErr error
	}

	tests := []test{
		{
			name: "should return error given invalid user_id",
			args: args{
				ctx: context.Background(),
				product: model.Product{
					Name:        "Some product name",
					Description: "lorem ipsum dolor sit amet",
					Price:       4500,
					Category:    "test",
					EanCode:     "1234567",
					UserId:      "12",
				},
			},
			wantErr: errors.New("Key: 'Product.UserId' Error:Field validation for 'UserId' failed on the 'uuid' tag"),
			mock: func(tt test) test {
				return tt
			},
		},
		{
			name: "should return no error",
			args: args{
				ctx: context.Background(),
				product: model.Product{
					Name:        "Some product name",
					Description: "lorem ipsum dolor sit amet",
					Price:       4500,
					Category:    "test",
					EanCode:     "1234567",
					UserId:      uuid.NewString(),
				},
			},
			wantErr: nil,
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)

				service.EXPECT().
					Add(tt.args.ctx, tt.args.product).
					Return(nil).
					Times(1)

				tt.fields.Service = service

				return tt
			},
		},
		{
			name: "should return error if service return error",
			args: args{
				ctx: context.Background(),
				product: model.Product{
					Name:        "Some product name",
					Description: "lorem ipsum dolor sit amet",
					Price:       4500,
					Category:    "test",
					EanCode:     "1234567",
					UserId:      uuid.NewString(),
				},
			},
			wantErr: errors.New("some error"),
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)

				service.EXPECT().
					Add(tt.args.ctx, tt.args.product).
					Return(errors.New("some error")).
					Times(1)

				tt.fields.Service = service

				return tt
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt = tt.mock(tt)

			handler := &handler.Handler{
				Service:   tt.fields.Service,
				Validator: validator.New(validator.WithRequiredStructEnabled()),
			}

			err := handler.Add(tt.args.ctx, tt.args.product)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

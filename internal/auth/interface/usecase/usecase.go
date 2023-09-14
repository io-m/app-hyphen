package auth_usecase_interface

import (
	"context"

	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

type IAuthUsecase interface {
	Register(ctx context.Context, customerRequest *customer_objects.CustomerRequest) (*customer_objects.CustomerResponse, error)
	Login(ctx context.Context, customerRequest *customer_objects.LoginCustomerRequest) (*customer.Customer, error)
	OAuth(ctx context.Context, customerRequest *customer_objects.LoginCustomerRequest) (*customer_objects.CustomerResponse, error)
}

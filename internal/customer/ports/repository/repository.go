package customer_repository

import (
	"context"

	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	"github.com/io-m/app-hyphen/pkg/types/tokens"
)

type ICustomerRepository interface {
	ICustomerCrud
	tokens.ITokens
}

type ICustomerCrud interface {
	CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error)
	GetAllCustomers(ctx context.Context) ([]*customer.Customer, error)
	GetCustomerById(ctx context.Context, customerId string) (*customer.Customer, error)
	UpdateCustomer(ctx context.Context, customerId string, customerRequest *customer_objects.CustomerRequestOptional) (*customer.Customer, error)
	DeleteCustomerById(ctx context.Context, customerId string) (string, error)
}

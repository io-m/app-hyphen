package customer_incoming

import (
	"context"

	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	"github.com/io-m/app-hyphen/pkg/types/tokens"
)

type ICustomerIncoming interface {
	ICustomerUsecase
	tokens.IAuthenticator
	tokens.ITokens
}

type ICustomerUsecase interface {
	ValidateCustomerPassword(customerRequest *customer_objects.CustomerRequest) error
	CreateCustomer(ctx context.Context, customerRequest *customer_objects.CustomerRequest) (*customer.Customer, error)
	GetAllCustomers(ctx context.Context) ([]*customer.Customer, error)
	GetCustomerById(ctx context.Context, customerId string) (*customer.Customer, error)
	UpdateCustomer(ctx context.Context, customerId string, customerRequest *customer_objects.CustomerRequest) (*customer.Customer, error)
	DeleteCustomerById(ctx context.Context, customerId string) (string, error)
}

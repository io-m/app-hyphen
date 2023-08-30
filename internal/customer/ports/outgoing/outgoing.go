package customer_outgoing

import (
	"context"

	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

type ICustomerOutgoing interface {
	CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error)
	GetAllCustomers(ctx context.Context) ([]*customer.Customer, error)
	GetCustomerById(ctx context.Context, customerId string) (*customer.Customer, error)
	UpdateCustomer(ctx context.Context, customerId string, customerRequest *customer_objects.CustomerRequest) (*customer.Customer, error)
	DeleteCustomerById(ctx context.Context, customerId string) (string, error)
	SaveRefreshToken(ctx context.Context, refreshToken string) (string, error)
	RetrieveRefreshToken(ctx context.Context, refreshToken string) (string, error)
}

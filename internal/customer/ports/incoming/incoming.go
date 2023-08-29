package customer_incoming

import (
	"context"

	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

type ICustomerIngoing interface {
	CreateCustomer(ctx context.Context, customerRequest *customer_objects.CustomerRequest) (any, error)
	GetAllCustomers(ctx context.Context) (any, error)
	GetCustomerById(ctx context.Context, bookId string) (any, error)
	UpdateCustomer(ctx context.Context, bookId string, customerRequest *customer_objects.CustomerRequest) (any, error)
	DeleteCustomerById(ctx context.Context, bookId string) (any, error)
}

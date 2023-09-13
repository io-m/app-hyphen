package customer_repository

import (
	"context"

	"github.com/google/uuid"
	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

type ICustomerRepository interface {
	ICustomerCrud
}

type ICustomerCrud interface {
	FindAllCustomers(ctx context.Context) ([]*customer.Customer, error)
	FindCustomerById(ctx context.Context, customerId uuid.UUID) (*customer.Customer, error)
	FindCustomerByEmail(ctx context.Context, email string) (*customer.Customer, error)
	CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error)
	UpdateCustomerById(ctx context.Context, customerId uuid.UUID, customerRequest *customer_objects.CustomerRequestOptional) (*customer.Customer, error)
	DeleteCustomerById(ctx context.Context, customerId uuid.UUID) (bool, error)
}

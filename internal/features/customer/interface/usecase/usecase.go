package customer_usecase

import (
	"context"

	"github.com/google/uuid"
	customer "github.com/io-m/app-hyphen/internal/features/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/features/customer/domain/objects"
)

type ICustomerUsecase interface {
	GetCustomerWithEmail(ctx context.Context, customerId string) (*customer.Customer, error)
	GetCustomerWithId(ctx context.Context, customerId uuid.UUID) (*customer_objects.CustomerResponse, error)
	UpdateCustomerWithId(ctx context.Context, customerId uuid.UUID, customerRequestOptional *customer_objects.CustomerRequestOptional) (*customer_objects.CustomerResponse, error)
	DeleteCustomerWithId(ctx context.Context, customerId uuid.UUID) (bool, error)
}

package customer_logic

import (
	"context"

	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

func (cl *customerLogic) CreateCustomer(ctx context.Context, customerRequest *customer_objects.CustomerRequest) (*customer.Customer, error) {
	customer := customer_objects.MapCustomerRequestToCustomer(customerRequest)
	return cl.customerOutgoing.CreateCustomer(ctx, customer)
}

package customer_logic

import (
	"context"

	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

func (cl *customerLogic) UpdateCustomer(ctx context.Context, customerId string, customerRequest *customer_objects.CustomerRequestOptional) (*customer.Customer, error) {
	return cl.customerOutgoing.UpdateCustomer(ctx, customerId, customerRequest)
}

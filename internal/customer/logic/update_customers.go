package customer_logic

import (
	"context"

	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

func (cl *customerLogic) UpdateCustomer(ctx context.Context, bookId string, customerRequest *customer_objects.CustomerRequest) (*customer.Customer, error) {
	return nil, nil
}

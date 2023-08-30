package customer_logic

import (
	"context"

	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
)

func (cl *customerLogic) GetAllCustomers(ctx context.Context) ([]*customer.Customer, error) {
	return nil, nil
}
func (cl *customerLogic) GetCustomerById(ctx context.Context, bookId string) (*customer.Customer, error) {
	return nil, nil
}

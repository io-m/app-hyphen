package customer_logic

import (
	"context"

	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	customer_outgoing "github.com/io-m/app-hyphen/internal/customer/ports/outgoing"
)

type customerLogic struct {
	customerOutgoing customer_outgoing.ICustomerOutgoing
}

func NewCustomerLogic(customerOutgoing customer_outgoing.ICustomerOutgoing) *customerLogic {
	return &customerLogic{
		customerOutgoing: customerOutgoing,
	}
}

func (cl *customerLogic) CreateCustomer(ctx context.Context, customerRequest *customer_objects.CustomerRequest) (any, error) {
	return nil, nil
}
func (cl *customerLogic) GetAllCustomers(ctx context.Context) (any, error) {
	return nil, nil
}
func (cl *customerLogic) GetCustomerById(ctx context.Context, bookId string) (any, error) {
	return nil, nil
}
func (cl *customerLogic) UpdateCustomer(ctx context.Context, bookId string, customerRequest *customer_objects.CustomerRequest) (any, error) {
	return nil, nil
}
func (cl *customerLogic) DeleteCustomerById(ctx context.Context, bookId string) (any, error) {
	return nil, nil
}

package customer_logic

import (
	"context"

	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	customer_outgoing "github.com/io-m/app-hyphen/internal/customer/ports/outgoing"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

type customerLogic struct {
	customerOutgoing customer_outgoing.ICustomerOutgoing
}

func NewCustomerLogic(customerOutgoing customer_outgoing.ICustomerOutgoing) *customerLogic {
	return &customerLogic{
		customerOutgoing: customerOutgoing,
	}
}

func (cl *customerLogic) ValidateCustomerPassword(customerRequest *customer_objects.CustomerRequest) error {
	if err := helpers.ValidatePassword(customerRequest.Password); err != nil {
		return err
	}
	return nil
}

func (cl *customerLogic) CreateCustomer(ctx context.Context, customerRequest *customer_objects.CustomerRequest) (*customer.Customer, error) {
	customer := customer_objects.MapCustomerRequestToCustomer(customerRequest)
	return cl.customerOutgoing.CreateCustomer(ctx, customer)
}
func (cl *customerLogic) GetAllCustomers(ctx context.Context) ([]*customer.Customer, error) {
	return nil, nil
}
func (cl *customerLogic) GetCustomerById(ctx context.Context, bookId string) (*customer.Customer, error) {
	return nil, nil
}
func (cl *customerLogic) UpdateCustomer(ctx context.Context, bookId string, customerRequest *customer_objects.CustomerRequest) (*customer.Customer, error) {
	return nil, nil
}
func (cl *customerLogic) DeleteCustomerById(ctx context.Context, bookId string) (string, error) {
	return "", nil
}

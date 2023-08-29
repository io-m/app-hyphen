package customer_arango_adapter

import (
	"context"

	"github.com/arangodb/go-driver"
	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

type arangoCustomerDB struct {
	arango driver.Database
}

func NewArangoCustomerDB(arangoDriver driver.Database) *arangoCustomerDB {
	return &arangoCustomerDB{
		arango: arangoDriver,
	}
}

// TODO: implement ICustomerOutgoing interface for Arango
func (arango *arangoCustomerDB) CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error) {
	return nil, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (arango *arangoCustomerDB) GetAllCustomers(ctx context.Context) ([]*customer.Customer, error) {
	return nil, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (arango *arangoCustomerDB) GetCustomerById(ctx context.Context, customerId string) (*customer.Customer, error) {
	return nil, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (arango *arangoCustomerDB) UpdateCustomer(ctx context.Context, customerId string, customerRequest *customer_objects.CustomerRequest) (*customer.Customer, error) {
	return nil, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (arango *arangoCustomerDB) DeleteCustomerById(ctx context.Context, customerId string) (string, error) {
	return "", nil
}

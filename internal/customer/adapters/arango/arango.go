package customer_arango_adapter

import (
	"context"
	"fmt"
	"log"

	"github.com/arangodb/go-driver"
	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

type arangoCustomerDB struct {
	driver driver.Database
}

func NewArangoCustomerDB(arangoDriver driver.Database) *arangoCustomerDB {
	return &arangoCustomerDB{
		driver: arangoDriver,
	}
}

// TODO: implement ICustomerOutgoing interface for Arango
func (arango *arangoCustomerDB) CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error) {
	// Perform AQL queries
	// Query to create customer
	log.Println("IN ARANGO STORE ---> ", customer)
	query := "INSERT @customer INTO customers"
	bindVars := map[string]interface{}{
		"customer": customer,
	}
	_, err := arango.driver.Query(ctx, query, bindVars)
	if err != nil {
		return nil, fmt.Errorf("failed to create Customer: %w", err)
	} else {
		fmt.Println("Successfully created customer.")
	}
	return customer, nil
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

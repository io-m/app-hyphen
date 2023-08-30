package customer_db_adapter

import (
	"context"
	"fmt"
	"log"

	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

func (db *customerOutgoing) CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error) {
	// Perform AQL queries
	// Query to create customer
	log.Println("IN db STORE ---> ", customer)
	query := "INSERT @customer INTO customers"
	bindVars := map[string]interface{}{
		"customer": customer,
	}

	_, err := db.arango.Query(ctx, query, bindVars)
	if err != nil {
		return nil, fmt.Errorf("failed to create Customer: %w", err)
	} else {
		fmt.Println("Successfully created customer.")
	}
	return customer, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (db *customerOutgoing) GetAllCustomers(ctx context.Context) ([]*customer.Customer, error) {
	return nil, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (db *customerOutgoing) GetCustomerById(ctx context.Context, customerId string) (*customer.Customer, error) {
	return nil, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (db *customerOutgoing) UpdateCustomer(ctx context.Context, customerId string, customerRequest *customer_objects.CustomerRequest) (*customer.Customer, error) {
	return nil, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (db *customerOutgoing) DeleteCustomerById(ctx context.Context, customerId string) (string, error) {
	return "", nil
}

package customer_db_adapter

import (
	"context"
	"fmt"

	"github.com/io-m/app-hyphen/internal/customer/adapters/database/queries"
	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

func (db *customerRepository) CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error) {
	// Perform AQL queries
	bindVars := map[string]any{
		"customer": customer,
	}

	_, err := db.arango.Query(ctx, queries.CREATE_CUSTOMER_QUERY, bindVars)
	if err != nil {
		return nil, fmt.Errorf("failed to create Customer: %w", err)
	} else {
		fmt.Println("Successfully created customer.")
	}
	return customer, nil
}

// TODO: implement ICustomerRepository interface for Arango
func (db *customerRepository) GetAllCustomers(ctx context.Context) ([]*customer.Customer, error) {
	return nil, nil
}

func (db *customerRepository) GetCustomerById(ctx context.Context, customerId string) (*customer.Customer, error) {
	bindVars := map[string]interface{}{
		"customerId": customerId,
	}
	cursor, err := db.arango.Query(ctx, queries.GET_CUSTOMER_BY_ID_QUERY, bindVars)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer cursor.Close()

	customer, err := helpers.ReadSingleDocument[customer.Customer](ctx, cursor)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (db *customerRepository) UpdateCustomer(ctx context.Context, customerId string, customerRequest *customer_objects.CustomerRequestOptional) (*customer.Customer, error) {
	updateQuery, bindVars := queries.BuildUpdateQueryAndVars(customerRequest, customerId)
	cursor, err := db.arango.Query(ctx, updateQuery, bindVars)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer cursor.Close()

	customer, err := helpers.ReadSingleDocument[customer.Customer](ctx, cursor)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (db *customerRepository) DeleteCustomerById(ctx context.Context, customerId string) (string, error) {
	return "", nil
}

package customer_arango_adapter

import (
	"context"
	"fmt"
	"log"

	"github.com/arangodb/go-driver"
	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

type customerDB struct {
	arango driver.Database
	redis  any
}

func NewCustomerDB(arangoDriver driver.Database, redisClient any) *customerDB {
	return &customerDB{
		arango: arangoDriver,
		redis:  redisClient,
	}
}

// TODO: implement ICustomerOutgoing interface for Arango
func (db *customerDB) CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error) {
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
func (db *customerDB) GetAllCustomers(ctx context.Context) ([]*customer.Customer, error) {
	return nil, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (db *customerDB) GetCustomerById(ctx context.Context, customerId string) (*customer.Customer, error) {
	return nil, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (db *customerDB) UpdateCustomer(ctx context.Context, customerId string, customerRequest *customer_objects.CustomerRequest) (*customer.Customer, error) {
	return nil, nil
}

// TODO: implement ICustomerOutgoing interface for Arango
func (db *customerDB) DeleteCustomerById(ctx context.Context, customerId string) (string, error) {
	return "", nil
}

// TODO: implement ICustomerOutgoing interface for Redis
func (db *customerDB) SaveRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	return "", nil
}

// TODO: implement ICustomerOutgoing interface for Redis
func (db *customerDB) RetrieveRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	return "", nil
}

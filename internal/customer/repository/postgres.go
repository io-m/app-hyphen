package customer_db_adapter

import (
	"context"
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

//go:embed queries/find_customer_by_id.sql
var findCustomerById string

func (db *customerRepository) CreateCustomer(ctx context.Context, customer *customer.Customer) (*customer.Customer, error) {

	return customer, nil
}

func (db *customerRepository) FindAllCustomers(ctx context.Context) ([]*customer.Customer, error) {
	return nil, nil
}

func (db *customerRepository) FindCustomerByEmail(ctx context.Context, email string) (*customer.Customer, error) {

	return nil, nil
}

func (db *customerRepository) FindCustomerById(ctx context.Context, customerId uuid.UUID) (*customer.Customer, error) {
	var customer customer.Customer
	err := db.postgres.Get(&customer, findCustomerById, customerId)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (db *customerRepository) UpdateCustomerById(ctx context.Context, customerId uuid.UUID, customerRequest *customer_objects.CustomerRequestOptional) (*customer.Customer, error) {
	query, args, err := buildUpdateSQL(customerId, customerRequest)
	if err != nil {
		return nil, err
	}

	_, err = db.postgres.NamedExec(query, args)
	if err != nil {
		return nil, err
	}

	updatedCustomer, err := db.FindCustomerById(ctx, customerId)
	if err != nil {
		return nil, err
	}
	return updatedCustomer, nil
}

// TODO: implement
func (db *customerRepository) DeleteCustomerById(ctx context.Context, customerId uuid.UUID) (bool, error) {
	return false, nil
}

func buildUpdateSQL(customerId uuid.UUID, customerRequest *customer_objects.CustomerRequestOptional) (string, map[string]interface{}, error) {
	setParts := []string{}
	args := map[string]interface{}{
		"id": customerId,
	}

	if customerRequest.FirstName != nil {
		setParts = append(setParts, "first_name = :first_name")
		args["first_name"] = *customerRequest.FirstName
	}

	if customerRequest.LastName != nil {
		setParts = append(setParts, "last_name = :last_name")
		args["last_name"] = *customerRequest.LastName
	}

	if customerRequest.Email != nil {
		setParts = append(setParts, "email = :email")
		args["email"] = *customerRequest.Email
	}

	// For password updates, you should handle hashing and verification outside this function.
	if customerRequest.NewPassword != nil {
		setParts = append(setParts, "password = :password")
		args["password"] = *customerRequest.NewPassword
	}

	// Add any other fields as necessary...

	if len(setParts) == 0 {
		return "", nil, fmt.Errorf("no fields to update")
	}

	// Always update the "updated_at" field.
	args["updated_at"] = time.Now().UTC().Format(time.RFC3339)
	setParts = append(setParts, "updated_at = :updated_at")

	sql := fmt.Sprintf("UPDATE customers SET %s WHERE id = :id", strings.Join(setParts, ", "))
	return sql, args, nil
}

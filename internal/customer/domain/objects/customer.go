package customer_objects

import (
	"time"

	"github.com/google/uuid"
	address "github.com/io-m/app-hyphen/internal/address/domain/entity"
	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	"github.com/io-m/app-hyphen/internal/tokens"
)

type CustomerRequest struct {
	FirstName string                    `json:"first_name"`
	LastName  string                    `json:"last_name"`
	Email     string                    `json:"email"`
	Password  string                    `json:"password"`
	Address   address.Address           `json:"address"`
	Role      tokens.AuthorizationLevel `json:"role"`
}

type CustomerRequestOptional struct {
	FirstName   *string                    `json:"first_name"`
	LastName    *string                    `json:"last_name"`
	Email       *string                    `json:"email"`
	OldPassword *string                    `json:"old_password"`
	NewPassword *string                    `json:"new_password"`
	Address     *address.Address           `json:"address"`
	Role        *tokens.AuthorizationLevel `json:"role"`
	UpdatedAt   string                     `json:"updated_at,omitempty"`
}

type LoginCustomerRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginCustomerResponse struct {
	AccessToken          string            `json:"access_token"`
	AccessTokenExpiresAt time.Time         `json:"access_token_expires_at"`
	Customer             customer.Customer `json:"customer"`
}

type CustomerResponse struct {
	ID        string                    `json:"id"`
	FirstName string                    `json:"first_name,omitempty"`
	LastName  string                    `json:"last_name,omitempty"`
	Email     string                    `json:"email,omitempty"`
	Address   address.Address           `json:"address,omitempty"`
	Role      tokens.AuthorizationLevel `json:"role,omitempty"`
	CreatedAt string                    `json:"created_at,omitempty"`
	UpdatedAt *string                   `json:"updated_at,omitempty"`
}

// MapCreateCustomerRequestToCustomer receives CustomerRequest and makes full Customer out of it in order to save it into DB
func MapCustomerRequestToCustomer(customerRequest *CustomerRequest) *customer.Customer {
	now := time.Now().Format(time.RFC3339)
	return &customer.Customer{
		ID:        uuid.NewString(),
		FirstName: customerRequest.FirstName,
		LastName:  customerRequest.LastName,
		Password:  customerRequest.Password,
		Email:     customerRequest.Email,
		Address:   customerRequest.Address,
		Role:      customerRequest.Role,
		CreatedAt: now,
		UpdatedAt: &now,
	}
}

// MapCreateCustomerRequestToCustomer receives CustomerRequest and makes full Customer out of it in order to save it into DB
func MapCustomerToCustomerResponse(customer *customer.Customer) *CustomerResponse {
	return &CustomerResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Address:   customer.Address,
		Role:      customer.Role,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}

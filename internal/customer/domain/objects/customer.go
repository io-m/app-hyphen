package customer_objects

import (
	"time"

	address "github.com/io-m/app-hyphen/internal/address/domain/entity"
	address_objects "github.com/io-m/app-hyphen/internal/address/domain/objects"
	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	"github.com/io-m/app-hyphen/internal/tokens"

	"github.com/google/uuid"
)

type CustomerRequest struct {
	FirstName string                          `json:"first_name"`
	LastName  string                          `json:"last_name"`
	Email     string                          `json:"email"`
	Password  string                          `json:"password"`
	Address   *address_objects.AddressRequest `json:"address"`
	Role      tokens.AuthorizationLevel       `json:"role"`
}

type CustomerRequestOptional struct {
	FirstName   *string                         `json:"first_name,omitempty"`
	LastName    *string                         `json:"last_name,omitempty"`
	Email       *string                         `json:"email,omitempty"`
	OldPassword *string                         `json:"old_password,omitempty"`
	NewPassword *string                         `json:"new_password,omitempty"`
	Address     *address_objects.AddressRequest `json:"address,omitempty"`
	Role        *tokens.AuthorizationLevel      `json:"role,omitempty"`
	UpdatedAt   *string                         `json:"updated_at,omitempty"`
}

type LoginCustomerRequest struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type LoginCustomerResponse struct {
	AccessToken          string            `json:"access_token"`
	AccessTokenExpiresAt time.Time         `json:"access_token_expires_at"`
	Customer             customer.Customer `json:"customer"`
}

type CustomerResponse struct {
	Id        uuid.UUID                 `json:"id"`
	FirstName string                    `json:"first_name,omitempty"`
	LastName  string                    `json:"last_name,omitempty"`
	Email     string                    `json:"email,omitempty"`
	Address   *address.Address          `json:"address,omitempty"`
	Role      tokens.AuthorizationLevel `json:"role,omitempty"`
	CreatedAt *string                   `json:"created_at,omitempty"`
	UpdatedAt *string                   `json:"updated_at,omitempty"`
}

// MapCreateCustomerRequestToCustomer receives CustomerRequest and makes full Customer out of it in order to save it into DB
func MapCustomerRequestToCustomer(customerRequest *CustomerRequest, address *address.Address) *customer.Customer {
	return &customer.Customer{
		Id:        uuid.New(),
		FirstName: customerRequest.FirstName,
		LastName:  customerRequest.LastName,
		Password:  customerRequest.Password,
		Email:     customerRequest.Email,
		Address:   address,
	}
}

// MapCreateCustomerRequestToCustomer receives CustomerRequest and makes full Customer out of it in order to save it into DB
func MapCustomerToCustomerResponse(customer *customer.Customer) *CustomerResponse {
	return &CustomerResponse{
		Id:        customer.Id,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Address:   customer.Address,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}

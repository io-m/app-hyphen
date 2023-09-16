package person_objects

import (
	"time"

	address "github.com/io-m/app-hyphen/internal/features/address/domain/entity"
	address_objects "github.com/io-m/app-hyphen/internal/features/address/domain/objects"
	person "github.com/io-m/app-hyphen/internal/features/person/domain/entity"
	"github.com/io-m/app-hyphen/internal/tokens"

	"github.com/google/uuid"
)

type PersonRequest struct {
	FirstName string                          `json:"first_name"`
	LastName  string                          `json:"last_name"`
	Email     string                          `json:"email"`
	Password  string                          `json:"password"`
	Address   *address_objects.AddressRequest `json:"address"`
	Role      tokens.AuthorizationLevel       `json:"role"`
}

type PersonRequestOptional struct {
	FirstName   *string                         `json:"first_name,omitempty"`
	LastName    *string                         `json:"last_name,omitempty"`
	Email       *string                         `json:"email,omitempty"`
	OldPassword *string                         `json:"old_password,omitempty"`
	NewPassword *string                         `json:"new_password,omitempty"`
	Address     *address_objects.AddressRequest `json:"address,omitempty"`
	Role        *tokens.AuthorizationLevel      `json:"role,omitempty"`
	UpdatedAt   *string                         `json:"updated_at,omitempty"`
}

type LoginPersonRequest struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type LoginPersonResponse struct {
	AccessToken          string        `json:"access_token"`
	AccessTokenExpiresAt time.Time     `json:"access_token_expires_at"`
	Person               person.Person `json:"person"`
}

type PersonResponse struct {
	Id        uuid.UUID                 `json:"id"`
	FirstName string                    `json:"first_name,omitempty"`
	LastName  string                    `json:"last_name,omitempty"`
	Email     string                    `json:"email,omitempty"`
	Address   *address.Address          `json:"address,omitempty"`
	Role      tokens.AuthorizationLevel `json:"role,omitempty"`
	CreatedAt *string                   `json:"created_at,omitempty"`
	UpdatedAt *string                   `json:"updated_at,omitempty"`
}

// MapPersonRequestToPerson receives PersonRequest and makes full Person out of it in order to save it into DB
func MapPersonRequestToPerson(personRequest *PersonRequest, address *address.Address) *person.Person {
	return &person.Person{
		Id:        uuid.New(),
		FirstName: personRequest.FirstName,
		LastName:  personRequest.LastName,
		Password:  personRequest.Password,
		Email:     personRequest.Email,
		Address:   address,
	}
}

// MapPersonToPersonResponse receives PersonRequest and makes full Person out of it in order to save it into DB
func MapPersonToPersonResponse(person *person.Person) *PersonResponse {
	return &PersonResponse{
		Id:        person.Id,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Email:     person.Email,
		Address:   person.Address,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
	}
}

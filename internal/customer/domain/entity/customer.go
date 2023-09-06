package customer

import (
	address "github.com/io-m/app-hyphen/internal/address/domain/entity"
	"github.com/io-m/app-hyphen/internal/tokens"
)

type Customer struct {
	ID        string                    `json:"id"`
	FirstName string                    `json:"first_name"`
	LastName  string                    `json:"last_name"`
	Email     string                    `json:"email"`
	Password  string                    `json:"password"`
	Address   address.Address           `json:"address"`
	Role      tokens.AuthorizationLevel `json:"role"`
	CreatedAt string                    `json:"created_at"`
	UpdatedAt *string                   `json:"updated_at,omitempty"`
}

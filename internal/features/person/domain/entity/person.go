package person

import (
	"github.com/google/uuid"
	"github.com/io-m/app-hyphen/internal/shared/types"
)

type Person struct {
	Id        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Address   *types.Address
	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at"`
}

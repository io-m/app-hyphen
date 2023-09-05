package customer_logic

import (
	"github.com/io-m/app-hyphen/pkg/helpers"
)

func (cl *customerLogic) ValidateCustomerPassword(customerPassword string) error {
	if err := helpers.ValidatePassword(customerPassword); err != nil {
		return err
	}
	return nil
}

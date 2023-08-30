package customer_logic

import (
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

func (cl *customerLogic) ValidateCustomerPassword(customerRequest *customer_objects.CustomerRequest) error {
	if err := helpers.ValidatePassword(customerRequest.Password); err != nil {
		return err
	}
	return nil
}

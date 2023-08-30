package customer_logic

import (
	customer_outgoing "github.com/io-m/app-hyphen/internal/customer/ports/outgoing"
)

type customerLogic struct {
	customerOutgoing customer_outgoing.ICustomerOutgoing
}

func NewCustomerLogic(customerOutgoing customer_outgoing.ICustomerOutgoing) *customerLogic {
	return &customerLogic{
		customerOutgoing: customerOutgoing,
	}
}

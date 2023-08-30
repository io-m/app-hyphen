package customer_logic

import (
	customer_outgoing "github.com/io-m/app-hyphen/internal/customer/ports/outgoing"
	"github.com/io-m/app-hyphen/pkg/types/tokens"
)

type customerLogic struct {
	customerOutgoing customer_outgoing.ICustomerOutgoing
	authenticator    tokens.IAuthenticator
}

func NewCustomerLogic(customerOutgoing customer_outgoing.ICustomerOutgoing, authenticator tokens.IAuthenticator) *customerLogic {
	return &customerLogic{
		customerOutgoing: customerOutgoing,
		authenticator:    authenticator,
	}
}

package customer_handler

import (
	"github.com/io-m/app-hyphen/pkg/types"
)

type ICustomerHandler interface {
	types.ICrudHandler
}

package customer_handler

import (
	"net/http"

	"github.com/io-m/app-hyphen/pkg/types"
)

type ICustomerHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	types.ICrudHandler
}

package customer_http_adapter

import (
	"net/http"

	customer_incoming "github.com/io-m/app-hyphen/internal/customer/ports/incoming"
)

type CustomerRESTHandler struct {
	customerIncoming customer_incoming.ICustomerIngoing
}

func NewCustomerRESTHandler(customerIncoming customer_incoming.ICustomerIngoing) *CustomerRESTHandler {
	return &CustomerRESTHandler{
		customerIncoming: customerIncoming,
	}
}

func (c *CustomerRESTHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {

}
func (c *CustomerRESTHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

}
func (c *CustomerRESTHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {

}
func (c *CustomerRESTHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {

}
func (c *CustomerRESTHandler) DeleteCustomerById(w http.ResponseWriter, r *http.Request) {

}

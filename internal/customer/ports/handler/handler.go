package customer_handler

import (
	"net/http"
)

type ICustomerHandler interface {
	LoginCustomer(w http.ResponseWriter, r *http.Request)
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	GetAllCustomers(w http.ResponseWriter, r *http.Request)
	GetCustomerById(w http.ResponseWriter, r *http.Request)
	UpdateCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomerById(w http.ResponseWriter, r *http.Request)
}

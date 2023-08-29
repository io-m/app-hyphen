package customer_http_adapter

import (
	"fmt"
	"net/http"

	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	customer_incoming "github.com/io-m/app-hyphen/internal/customer/ports/incoming"
	"github.com/io-m/app-hyphen/pkg/helpers"
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
	customerRequest, err := helpers.DecodePayload[*customer_objects.CustomerRequest](w, r)
	// c, err := helpers.DecodePayload[*customer.Customer](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusInternalServerError)
		return
	}
	if err := c.customerIncoming.ValidateCustomerPassword(customerRequest); err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error in password: %w", err), http.StatusBadRequest)
		return
	}
	// customer, err := hyphen.useCases.CustomerUsecase.FindByEmail(c.Email)
	// if err == nil {
	// 	utils.ErrorResponse(w, fmt.Errorf("email %s already registered", customer.Email), http.StatusBadRequest)
	// 	return
	// }
	hashedPassword, err := helpers.HashPassword(customerRequest.Password)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not hash password: %w", err), http.StatusBadRequest)
		return
	}
	// customer.Role = entities.CUSTOMER
	customerRequest.Password = hashedPassword
	customer, err := c.customerIncoming.CreateCustomer(r.Context(), customerRequest)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not register this email: %w", err), http.StatusInternalServerError)
		return
	}
	helpers.SuccessResponse(w, customer_objects.MapCustomerToCustomerResponse(customer), "Customer successfully registered", http.StatusCreated)
}
func (c *CustomerRESTHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

}
func (c *CustomerRESTHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {

}
func (c *CustomerRESTHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {

}
func (c *CustomerRESTHandler) DeleteCustomerById(w http.ResponseWriter, r *http.Request) {

}

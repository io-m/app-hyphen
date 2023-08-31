package customer_http_adapter

import (
	"fmt"
	"net/http"

	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	customer_incoming "github.com/io-m/app-hyphen/internal/customer/ports/incoming"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/helpers"
	"github.com/io-m/app-hyphen/pkg/types/tokens"
)

type CustomerRESTHandler struct {
	customerIncoming customer_incoming.ICustomerIncoming
}

func NewCustomerRESTHandler(customerIncoming customer_incoming.ICustomerIncoming) *CustomerRESTHandler {
	return &CustomerRESTHandler{
		customerIncoming: customerIncoming,
	}
}

func (ch *CustomerRESTHandler) LoginCustomer(w http.ResponseWriter, r *http.Request) {
	c, err := helpers.DecodePayload[*customer_objects.LoginCustomerRequest](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusInternalServerError)
		return
	}
	if err := helpers.ValidatePassword(c.Password); err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error in password: %w", err), http.StatusBadRequest)
		return
	}
	customer, err := ch.customerIncoming.GetCustomerById(r.Context(), c.ID)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not find customer: %w", err), http.StatusNotFound)
		return
	}
	if err := helpers.CheckPassword(c.Password, customer.Password); err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("wrong password: %w", err), http.StatusBadRequest)
		return
	}
	claims, _ := tokens.NewClaims(customer.ID, constants.ACCESS_TOKEN_DURATION)

	accessToken, refreshToken, err := ch.customerIncoming.GenerateTokens(claims)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while generating tokens: %w", err), http.StatusInternalServerError)
		return
	}
	// Here we need to save refresh token in Redis
	if err := ch.customerIncoming.SaveRefreshToken(r.Context(), refreshToken); err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while saving refresh token: %w", err), http.StatusInternalServerError)
		return
	}
	w.Header().Add(constants.ACCESS_TOKEN_HEADER, accessToken)
	w.Header().Add(constants.REFRESH_TOKEN_HEADER, refreshToken)

	helpers.SuccessResponse(w, customer_objects.MapCustomerToCustomerResponse(customer), "Customer successfully logged in")
}

func (ch *CustomerRESTHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	customerRequest, err := helpers.DecodePayload[*customer_objects.CustomerRequest](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusInternalServerError)
		return
	}
	if err := ch.customerIncoming.ValidateCustomerPassword(customerRequest); err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error in password: %w", err), http.StatusBadRequest)
		return
	}
	hashedPassword, err := helpers.HashPassword(customerRequest.Password)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not hash password: %w", err), http.StatusBadRequest)
		return
	}
	// customer.Role = entities.CUSTOMER
	customerRequest.Password = hashedPassword
	customer, err := ch.customerIncoming.CreateCustomer(r.Context(), customerRequest)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not register this email: %w", err), http.StatusInternalServerError)
		return
	}
	helpers.SuccessResponse(w, customer_objects.MapCustomerToCustomerResponse(customer), "Customer successfully registered", http.StatusCreated)
}
func (ch *CustomerRESTHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	//? Do we even need this one
}
func (ch *CustomerRESTHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId := helpers.GetUrlParam(r, "id")
	customer, err := ch.customerIncoming.GetCustomerById(r.Context(), customerId)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not find customer with id %s: %w", customerId, err), http.StatusNotFound)
		return
	}
	helpers.SuccessResponse(w, customer_objects.MapCustomerToCustomerResponse(customer), "Customer found", http.StatusOK)
}
func (ch *CustomerRESTHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {

}
func (ch *CustomerRESTHandler) DeleteCustomerById(w http.ResponseWriter, r *http.Request) {

}

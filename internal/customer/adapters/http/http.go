package customer_http_adapter

import (
	"fmt"
	"net/http"

	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	customer_handler "github.com/io-m/app-hyphen/internal/customer/ports/handler"
	customer_usecase "github.com/io-m/app-hyphen/internal/customer/ports/usecase"
	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

type CustomerRESTHandler struct {
	customerUsecase customer_usecase.ICustomerUsecase
	authenticator   tokens.ITokens
}

func NewCustomerRESTHandler(customerUsecase customer_usecase.ICustomerUsecase, authenticator tokens.ITokens) customer_handler.ICustomerHandler {
	return &CustomerRESTHandler{
		customerUsecase: customerUsecase,
		authenticator:   authenticator,
	}
}

func (ch *CustomerRESTHandler) Login(w http.ResponseWriter, r *http.Request) {
	c, err := helpers.DecodePayload[*customer_objects.LoginCustomerRequest](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}
	customer, err := ch.customerUsecase.GetCustomerWithEmail(r.Context(), c.Email)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not find customer: %w", err), http.StatusNotFound)
		return
	}
	if err := helpers.CheckPassword(c.Password, customer.Password); err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("wrong password: %w", err), http.StatusBadRequest)
		return
	}
	claims, _ := tokens.NewClaims(customer.Id, constants.ACCESS_TOKEN_DURATION)

	accessToken, refreshToken, err := ch.authenticator.GenerateTokens(claims)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while generating tokens: %w", err), http.StatusInternalServerError)
		return
	}
	// Here we need to save refresh token in Redis
	if err := ch.authenticator.SaveRefreshToken(r.Context(), customer.Id, refreshToken); err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while saving refresh token: %w", err), http.StatusInternalServerError)
		return
	}
	w.Header().Add(constants.ACCESS_TOKEN_HEADER, accessToken)
	w.Header().Add(constants.REFRESH_TOKEN_HEADER, refreshToken)

	helpers.SuccessResponse(w, customer_objects.MapCustomerToCustomerResponse(customer), "Customer successfully logged in")
}

func (ch *CustomerRESTHandler) Create(w http.ResponseWriter, r *http.Request) {
	customerRequest, err := helpers.DecodePayload[*customer_objects.CustomerRequest](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}
	customerResponse, err := ch.customerUsecase.CreateCustomer(r.Context(), customerRequest)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not register: %w", err), http.StatusBadRequest)
		return
	}
	helpers.SuccessResponse(w, customerResponse, "Customer successfully registered", http.StatusCreated)
}

func (ch *CustomerRESTHandler) GetById(w http.ResponseWriter, r *http.Request) {
	customerId := helpers.GetUrlParam(r, "id")
	customerResponse, err := ch.customerUsecase.GetCustomerWithId(r.Context(), customerId)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not find customer with id %s: %w", customerId, err), http.StatusNotFound)
		return
	}
	helpers.SuccessResponse(w, customerResponse, "Customer found", http.StatusOK)
}
func (ch *CustomerRESTHandler) Update(w http.ResponseWriter, r *http.Request) {
	customerId := helpers.GetUrlParam(r, "id")
	customerRequest, err := helpers.DecodePayload[*customer_objects.CustomerRequestOptional](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}
	customerResponse, err := ch.customerUsecase.UpdateCustomerWithId(r.Context(), customerId, customerRequest)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not update customer with id %s: %w", customerId, err), http.StatusInternalServerError)
		return
	}
	helpers.SuccessResponse(w, customerResponse, "Customer successfully updated", http.StatusNoContent)
}
func (ch *CustomerRESTHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	customerId := helpers.GetUrlParam(r, "id")
	ok, err := ch.customerUsecase.DeleteCustomerWithId(r.Context(), customerId)
	if !ok || err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not delete customer with id %s: %w", customerId, err), http.StatusInternalServerError)
		return
	}
}

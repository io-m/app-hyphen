package customer_http_adapter

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	customer_handler "github.com/io-m/app-hyphen/internal/customer/ports/handler"
	customer_repository "github.com/io-m/app-hyphen/internal/customer/ports/repository"
	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

type CustomerRESTHandler struct {
	customerRepo  customer_repository.ICustomerRepository
	authenticator tokens.ITokens
}

func NewCustomerRESTHandler(customerRepo customer_repository.ICustomerRepository, authenticator tokens.ITokens) customer_handler.ICustomerHandler {
	return &CustomerRESTHandler{
		customerRepo:  customerRepo,
		authenticator: authenticator,
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
	customer, err := ch.customerRepo.GetCustomerById(r.Context(), c.ID)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not find customer: %w", err), http.StatusNotFound)
		return
	}
	if err := helpers.CheckPassword(c.Password, customer.Password); err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("wrong password: %w", err), http.StatusBadRequest)
		return
	}
	claims, _ := tokens.NewClaims(customer.ID, constants.ACCESS_TOKEN_DURATION)

	accessToken, refreshToken, err := ch.authenticator.GenerateTokens(claims)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while generating tokens: %w", err), http.StatusInternalServerError)
		return
	}
	// Here we need to save refresh token in Redis
	if err := ch.authenticator.SaveRefreshToken(r.Context(), customer.ID, refreshToken); err != nil {
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
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusInternalServerError)
		return
	}
	if err := helpers.ValidatePassword(customerRequest.Password); err != nil {
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
	customer, err := ch.customerRepo.CreateCustomer(r.Context(), customer_objects.MapCustomerRequestToCustomer(customerRequest))
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not register this email: %w", err), http.StatusInternalServerError)
		return
	}
	helpers.SuccessResponse(w, customer_objects.MapCustomerToCustomerResponse(customer), "Customer successfully registered", http.StatusCreated)
}
func (ch *CustomerRESTHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	//? Do we even need this one
}
func (ch *CustomerRESTHandler) GetById(w http.ResponseWriter, r *http.Request) {
	customerId := helpers.GetUrlParam(r, "id")
	customer, err := ch.customerRepo.GetCustomerById(r.Context(), customerId)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not find customer with id %s: %w", customerId, err), http.StatusNotFound)
		return
	}
	helpers.SuccessResponse(w, customer_objects.MapCustomerToCustomerResponse(customer), "Customer found", http.StatusOK)
}
func (ch *CustomerRESTHandler) Update(w http.ResponseWriter, r *http.Request) {
	customerId := helpers.GetUrlParam(r, "id")
	claims, ok := r.Context().Value(constants.CLAIMS).(*tokens.Claims)
	if !ok {
		helpers.ErrorResponse(w, errors.New("token verification issue:cannot extract claims from context"), http.StatusUnauthorized)
		return
	}
	log.Println("CUSTOMER ID ===> ", customerId)
	log.Println("CLAIM SUBJECT ID ===> ", claims.SubjectID)
	log.Println("CLAIM  ID ===> ", claims.ClaimID)
	if claims.SubjectID != customerId {
		helpers.ErrorResponse(w, errors.New("token verification issue: Not your profile"), http.StatusUnauthorized)
		return
	}
	customerRequest, err := helpers.DecodePayload[*customer_objects.CustomerRequestOptional](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}
	if customerRequest.OldPassword != nil {
		if err := helpers.ValidatePassword(*customerRequest.OldPassword); err != nil {
			helpers.ErrorResponse(w, fmt.Errorf("error in password: %w", err), http.StatusBadRequest)
			return
		}
	}
	if customerRequest.OldPassword != nil {
		hashedPassword, err := helpers.HashPassword(*customerRequest.NewPassword)
		if err != nil {
			helpers.ErrorResponse(w, fmt.Errorf("could not hash password: %w", err), http.StatusBadRequest)
			return
		}
		customerRequest.NewPassword = &hashedPassword
	}
	customer, err := ch.customerRepo.UpdateCustomer(r.Context(), customerId, customerRequest)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not register this email: %w", err), http.StatusInternalServerError)
		return
	}
	helpers.SuccessResponse(w, customer_objects.MapCustomerToCustomerResponse(customer), "Customer successfully updated", http.StatusNoContent)
}
func (ch *CustomerRESTHandler) DeleteById(w http.ResponseWriter, r *http.Request) {

}

package customer_http_adapter

import (
	"fmt"
	"net/http"

	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	customer_usecase "github.com/io-m/app-hyphen/internal/customer/interface/usecase"
	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

type customerHandler struct {
	customerUsecase customer_usecase.ICustomerUsecase
	protector       tokens.IProtector
}

func NewCustomerHandler(customerUsecase customer_usecase.ICustomerUsecase, protector tokens.IProtector) *customerHandler {
	return &customerHandler{
		customerUsecase: customerUsecase,
		protector:       protector,
	}
}

func (ch *customerHandler) GetById(w http.ResponseWriter, r *http.Request) {
	customerId := helpers.GetUrlParam(r, "id")
	customerResponse, err := ch.customerUsecase.GetCustomerWithId(r.Context(), customerId)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not find customer with id %s: %w", customerId, err), http.StatusNotFound)
		return
	}
	helpers.SuccessResponse(w, customerResponse, "Customer found", http.StatusOK)
}
func (ch *customerHandler) Update(w http.ResponseWriter, r *http.Request) {
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
func (ch *customerHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	customerId := helpers.GetUrlParam(r, "id")
	ok, err := ch.customerUsecase.DeleteCustomerWithId(r.Context(), customerId)
	if !ok || err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not delete customer with id %s: %w", customerId, err), http.StatusInternalServerError)
		return
	}
}

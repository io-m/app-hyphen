package customer_usecase_adapter

import (
	"context"
	"errors"

	"github.com/google/uuid"
	address_repository "github.com/io-m/app-hyphen/internal/address/interface/repository"
	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	customer_repository "github.com/io-m/app-hyphen/internal/customer/interface/repository"
	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

type customerUsecase struct {
	addressRepo   address_repository.IAddressRepository
	customerRepo  customer_repository.ICustomerRepository
	authenticator tokens.ITokens
}

func NewCustomerUsecase(addressRepo address_repository.IAddressRepository, customerRepo customer_repository.ICustomerRepository, authenticator tokens.ITokens) *customerUsecase {
	return &customerUsecase{
		addressRepo:   addressRepo,
		customerRepo:  customerRepo,
		authenticator: authenticator,
	}
}

/* ICustomerUsecase interface implementations */
func (cu *customerUsecase) GetCustomerWithEmail(ctx context.Context, email string) (*customer.Customer, error) {
	savedCustomer, err := cu.customerRepo.FindCustomerByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return savedCustomer, nil
}

func (cu *customerUsecase) GetCustomerWithId(ctx context.Context, customerId uuid.UUID) (*customer_objects.CustomerResponse, error) {
	savedCustomer, err := cu.customerRepo.FindCustomerById(ctx, customerId)
	if err != nil {
		return nil, err
	}
	return customer_objects.MapCustomerToCustomerResponse(savedCustomer), nil
}

func (cu *customerUsecase) UpdateCustomerWithId(ctx context.Context, customerId uuid.UUID, customerRequestOptional *customer_objects.CustomerRequestOptional) (*customer_objects.CustomerResponse, error) {
	if !helpers.IsUserAuthorized(ctx, customerId) {
		return nil, errors.New("not authorized")
	}
	if customerRequestOptional.OldPassword != nil {
		if err := helpers.ValidatePassword(*customerRequestOptional.OldPassword); err != nil {
			return nil, err
		}
	}
	if customerRequestOptional.OldPassword != nil {
		hashedPassword, err := helpers.HashPassword(*customerRequestOptional.NewPassword)
		if err != nil {
			return nil, err
		}
		customerRequestOptional.NewPassword = &hashedPassword
	}
	customer, err := cu.customerRepo.UpdateCustomerById(ctx, customerId, customerRequestOptional)
	if err != nil {
		return nil, err
	}
	return customer_objects.MapCustomerToCustomerResponse(customer), nil
}

func (cu *customerUsecase) DeleteCustomerWithId(ctx context.Context, customerId uuid.UUID) (bool, error) {
	return cu.customerRepo.DeleteCustomerById(ctx, customerId)
}

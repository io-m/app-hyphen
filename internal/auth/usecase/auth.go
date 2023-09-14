package auth_usecase

import (
	"context"
	"errors"
	"os"

	address_repository_interface "github.com/io-m/app-hyphen/internal/address/interface/repository"
	customer "github.com/io-m/app-hyphen/internal/customer/domain/entity"
	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
	customer_repository "github.com/io-m/app-hyphen/internal/customer/interface/repository"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

type authUsecase struct {
	addressRepository  address_repository_interface.IAddressRepository
	customerRepository customer_repository.ICustomerRepository
	accessTokenSecret  string
	refreshTokenSecret string
}

func NewAuthUsecase(addressRepository address_repository_interface.IAddressRepository, customerRepository customer_repository.ICustomerRepository) *authUsecase {
	return &authUsecase{
		addressRepository:  addressRepository,
		customerRepository: customerRepository,
		accessTokenSecret:  os.Getenv(constants.ACCESS_TOKEN_SECRET_KEY),
		refreshTokenSecret: os.Getenv(constants.REFRESH_TOKEN_SECRET_KEY),
	}
}

func (au *authUsecase) Register(ctx context.Context, customerRequest *customer_objects.CustomerRequest) (*customer_objects.CustomerResponse, error) {
	if err := helpers.ValidateName(customerRequest.FirstName); err != nil {
		return nil, errors.New("first name is invalid")
	}
	if err := helpers.ValidateName(customerRequest.LastName); err != nil {
		return nil, errors.New("last name is invalid")
	}
	if err := helpers.ValidatePassword(customerRequest.Password); err != nil {
		return nil, errors.New("password is invalid")
	}
	if err := helpers.ValidateEmail(customerRequest.Email); err != nil {
		return nil, errors.New("email is invalid")
	}
	hashedPassword, err := helpers.HashPassword(customerRequest.Password)
	if err != nil {
		return nil, err
	}
	customerRequest.Password = hashedPassword
	createdAddress, err := au.addressRepository.CreateAddress(ctx, customerRequest.Address)
	if err != nil {
		return nil, err
	}
	createdCustomer, err := au.customerRepository.CreateCustomer(ctx, customer_objects.MapCustomerRequestToCustomer(customerRequest, createdAddress))
	if err != nil {
		return nil, err
	}
	return customer_objects.MapCustomerToCustomerResponse(createdCustomer), nil
}

func (au *authUsecase) Login(ctx context.Context, customerRequest *customer_objects.LoginCustomerRequest) (*customer.Customer, error) {
	return nil, nil
}

func (au *authUsecase) OAuth(ctx context.Context, customerRequest *customer_objects.LoginCustomerRequest) (*customer_objects.CustomerResponse, error) {
	return nil, nil
}

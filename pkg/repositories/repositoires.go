package repositories

import (
	address_repository_interface "github.com/io-m/app-hyphen/internal/address/interface/repository"
	customer_repository_interface "github.com/io-m/app-hyphen/internal/customer/interface/repository"
)

type AppRepositories struct {
	AddressRepository  address_repository_interface.IAddressRepository
	CustomerRepository customer_repository_interface.ICustomerRepository
}

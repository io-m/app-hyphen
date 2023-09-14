package address_repository_interface

import (
	"context"

	address "github.com/io-m/app-hyphen/internal/address/domain/entity"
	address_objects "github.com/io-m/app-hyphen/internal/address/domain/objects"
)

type IAddressRepository interface {
	CreateAddress(ctx context.Context, addressRequest *address_objects.AddressRequest) (*address.Address, error)
	FindAllAddresses(ctx context.Context) ([]*address.Address, error)
	FindAddressById(ctx context.Context, addressId string) (*address.Address, error)
	UpdateAddress(ctx context.Context, addressId string, addressRequest *address_objects.AddressRequest) (*address.Address, error)
	DeleteAddressById(ctx context.Context, addressId string) (bool, error)
}

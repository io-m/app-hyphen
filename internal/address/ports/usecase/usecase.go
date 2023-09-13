package address_usecase

import (
	"context"

	address "github.com/io-m/app-hyphen/internal/address/domain/entity"
	address_objects "github.com/io-m/app-hyphen/internal/address/domain/objects"
)

type IAddressUsecase interface {
	GetAllAddresses(ctx context.Context, addressId string) ([]*address.Address, error)
	GetAddressWithId(ctx context.Context, addressId string) (*address.Address, error)
	CreateAddress(ctx context.Context, addressRequest *address_objects.AddressRequest) (*address.Address, error)
	UpdateAddressWithId(ctx context.Context, addressId string, addressRequest *address_objects.AddressRequest) (*address.Address, error)
	DeleteAddressWithId(ctx context.Context, addressId string) (bool, error)
}

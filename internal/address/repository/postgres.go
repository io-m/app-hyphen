package address_repository

import (
	"context"

	address "github.com/io-m/app-hyphen/internal/address/domain/entity"
	address_objects "github.com/io-m/app-hyphen/internal/address/domain/objects"
)

func (ar *addressRepository) CreateAddress(ctx context.Context, addressRequest *address_objects.AddressRequest) (*address.Address, error) {
	return nil, nil
}
func (ar *addressRepository) FindAllAddresses(ctx context.Context) ([]*address.Address, error) {
	return nil, nil
}
func (ar *addressRepository) FindAddressById(ctx context.Context, addressId string) (*address.Address, error) {
	return nil, nil
}
func (ar *addressRepository) UpdateAddress(ctx context.Context, addressId string, addressRequest *address_objects.AddressRequest) (*address.Address, error) {
	return nil, nil
}
func (ar *addressRepository) DeleteAddressById(ctx context.Context, addressId string) (bool, error) {
	return false, nil
}

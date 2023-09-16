package person_repository_interface

import (
	"context"

	"github.com/google/uuid"
	person "github.com/io-m/app-hyphen/internal/features/person/domain/entity"
	person_objects "github.com/io-m/app-hyphen/internal/features/person/domain/objects"
	"github.com/io-m/app-hyphen/internal/shared/types"
)

type IPersonRepository interface {
	FindAllPersons(ctx context.Context) ([]*person.Person, error)
	FindPersonById(ctx context.Context, personId uuid.UUID) (*person.Person, error)
	FindPersonByEmail(ctx context.Context, email string) (*person.Person, error)
	CreatePerson(ctx context.Context, person *person.Person) (*person.Person, error)
	UpdatePersonById(ctx context.Context, personId uuid.UUID, personRequest *person_objects.PersonRequestOptional) (*person.Person, error)
	DeletePersonById(ctx context.Context, personId uuid.UUID) (bool, error)

	CreateAddress(ctx context.Context, addressRequest *types.AddressRequest) (*types.Address, error)
}

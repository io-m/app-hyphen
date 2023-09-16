package person_usecase

import (
	"context"

	"github.com/google/uuid"
	person "github.com/io-m/app-hyphen/internal/features/person/domain/entity"
	person_objects "github.com/io-m/app-hyphen/internal/features/person/domain/objects"
)

type IPersonUsecase interface {
	GetPersonWithEmail(ctx context.Context, personId string) (*person.Person, error)
	GetPersonWithId(ctx context.Context, personId uuid.UUID) (*person_objects.PersonResponse, error)
	UpdatePersonWithId(ctx context.Context, personId uuid.UUID, personRequestOptional *person_objects.PersonRequestOptional) (*person_objects.PersonResponse, error)
	DeletePersonWithId(ctx context.Context, personId uuid.UUID) (bool, error)
}

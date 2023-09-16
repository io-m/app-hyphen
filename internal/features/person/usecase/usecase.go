package person_usecase_adapter

import (
	"context"
	"errors"

	"github.com/google/uuid"
	person "github.com/io-m/app-hyphen/internal/features/person/domain/entity"
	person_objects "github.com/io-m/app-hyphen/internal/features/person/domain/objects"
	person_repository_interface "github.com/io-m/app-hyphen/internal/features/person/interface/repository"
	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

type personUsecase struct {
	personRepo    person_repository_interface.IPersonRepository
	authenticator tokens.ITokens
}

func NewPersonUsecase(personRepo person_repository_interface.IPersonRepository, authenticator tokens.ITokens) *personUsecase {
	return &personUsecase{
		personRepo:    personRepo,
		authenticator: authenticator,
	}
}

/* IPersonUsecase interface implementations */
func (cu *personUsecase) GetPersonWithEmail(ctx context.Context, email string) (*person.Person, error) {
	savedPerson, err := cu.personRepo.FindPersonByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return savedPerson, nil
}

func (cu *personUsecase) GetPersonWithId(ctx context.Context, personId uuid.UUID) (*person_objects.PersonResponse, error) {
	savedPerson, err := cu.personRepo.FindPersonById(ctx, personId)
	if err != nil {
		return nil, err
	}
	return person_objects.MapPersonToPersonResponse(savedPerson), nil
}

func (cu *personUsecase) UpdatePersonWithId(ctx context.Context, personId uuid.UUID, personRequestOptional *person_objects.PersonRequestOptional) (*person_objects.PersonResponse, error) {
	if !helpers.IsUserAuthorized(ctx, personId) {
		return nil, errors.New("not authorized")
	}
	if personRequestOptional.OldPassword != nil {
		if err := helpers.ValidatePassword(*personRequestOptional.OldPassword); err != nil {
			return nil, err
		}
	}
	if personRequestOptional.OldPassword != nil {
		hashedPassword, err := helpers.HashPassword(*personRequestOptional.NewPassword)
		if err != nil {
			return nil, err
		}
		personRequestOptional.NewPassword = &hashedPassword
	}
	person, err := cu.personRepo.UpdatePersonById(ctx, personId, personRequestOptional)
	if err != nil {
		return nil, err
	}
	return person_objects.MapPersonToPersonResponse(person), nil
}

func (cu *personUsecase) DeletePersonWithId(ctx context.Context, personId uuid.UUID) (bool, error) {
	return cu.personRepo.DeletePersonById(ctx, personId)
}

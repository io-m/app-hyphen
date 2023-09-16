package auth_usecase

import (
	"context"
	"errors"
	"os"

	person "github.com/io-m/app-hyphen/internal/features/person/domain/entity"
	person_objects "github.com/io-m/app-hyphen/internal/features/person/domain/objects"
	person_repository_interface "github.com/io-m/app-hyphen/internal/features/person/interface/repository"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

type authUsecase struct {
	personRepository   person_repository_interface.IPersonRepository
	accessTokenSecret  string
	refreshTokenSecret string
}

func NewAuthUsecase(personRepository person_repository_interface.IPersonRepository) *authUsecase {
	return &authUsecase{
		personRepository:   personRepository,
		accessTokenSecret:  os.Getenv(constants.ACCESS_TOKEN_SECRET_KEY),
		refreshTokenSecret: os.Getenv(constants.REFRESH_TOKEN_SECRET_KEY),
	}
}

func (au *authUsecase) Register(ctx context.Context, personRequest *person_objects.PersonRequest) (*person_objects.PersonResponse, error) {
	if err := helpers.ValidateName(personRequest.FirstName); err != nil {
		return nil, errors.New("first name is invalid")
	}
	if err := helpers.ValidateName(personRequest.LastName); err != nil {
		return nil, errors.New("last name is invalid")
	}
	if err := helpers.ValidatePassword(personRequest.Password); err != nil {
		return nil, errors.New("password is invalid")
	}
	if err := helpers.ValidateEmail(personRequest.Email); err != nil {
		return nil, errors.New("email is invalid")
	}
	hashedPassword, err := helpers.HashPassword(personRequest.Password)
	if err != nil {
		return nil, err
	}
	personRequest.Password = hashedPassword
	createdAddress, err := au.personRepository.CreateAddress(ctx, personRequest.Address)
	if err != nil {
		return nil, err
	}
	createdPerson, err := au.personRepository.CreatePerson(ctx, person_objects.MapPersonRequestToPerson(personRequest, createdAddress))
	if err != nil {
		return nil, err
	}
	return person_objects.MapPersonToPersonResponse(createdPerson), nil
}

func (au *authUsecase) Login(ctx context.Context, personRequest *person_objects.LoginPersonRequest) (*person.Person, error) {
	return nil, nil
}

func (au *authUsecase) OAuth(ctx context.Context, personRequest *person_objects.LoginPersonRequest) (*person_objects.PersonResponse, error) {
	return nil, nil
}

package auth_usecase_interface

import (
	"context"

	person "github.com/io-m/app-hyphen/internal/features/person/domain/entity"
	person_objects "github.com/io-m/app-hyphen/internal/features/person/domain/objects"
)

type IAuthUsecase interface {
	Register(ctx context.Context, personRequest *person_objects.PersonRequest) (*person_objects.PersonResponse, error)
	Login(ctx context.Context, personRequest *person_objects.LoginPersonRequest) (*person.Person, error)
	OAuth(ctx context.Context, personRequest *person_objects.LoginPersonRequest) (*person_objects.PersonResponse, error)
}

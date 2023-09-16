package person_repository

import (
	"context"
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	person "github.com/io-m/app-hyphen/internal/features/person/domain/entity"
	person_objects "github.com/io-m/app-hyphen/internal/features/person/domain/objects"
	"github.com/io-m/app-hyphen/internal/shared/types"
)

//go:embed queries/find_person_by_id.sql
var findPersonById string

func (db *personRepository) CreatePerson(ctx context.Context, person *person.Person) (*person.Person, error) {

	return person, nil
}

func (db *personRepository) FindAllPersons(ctx context.Context) ([]*person.Person, error) {
	return nil, nil
}

func (db *personRepository) FindPersonByEmail(ctx context.Context, email string) (*person.Person, error) {

	return nil, nil
}

func (db *personRepository) FindPersonById(ctx context.Context, personId uuid.UUID) (*person.Person, error) {
	var person person.Person
	err := db.postgres.GetContext(ctx, &person, findPersonById, personId)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (db *personRepository) UpdatePersonById(ctx context.Context, personId uuid.UUID, personRequest *person_objects.PersonRequestOptional) (*person.Person, error) {
	query, args, err := buildUpdateSQL(personId, personRequest)
	if err != nil {
		return nil, err
	}

	_, err = db.postgres.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, err
	}

	updatedPerson, err := db.FindPersonById(ctx, personId)
	if err != nil {
		return nil, err
	}
	return updatedPerson, nil
}

// TODO: implement
func (db *personRepository) DeletePersonById(ctx context.Context, personId uuid.UUID) (bool, error) {
	return false, nil
}

func (db *personRepository) CreateAddress(ctx context.Context, addressRequest *types.AddressRequest) (*types.Address, error) {
	return nil, nil
}

func buildUpdateSQL(personId uuid.UUID, personRequest *person_objects.PersonRequestOptional) (string, map[string]interface{}, error) {
	setParts := []string{}
	args := map[string]interface{}{
		"id": personId,
	}

	if personRequest.FirstName != nil {
		setParts = append(setParts, "first_name = :first_name")
		args["first_name"] = *personRequest.FirstName
	}

	if personRequest.LastName != nil {
		setParts = append(setParts, "last_name = :last_name")
		args["last_name"] = *personRequest.LastName
	}

	if personRequest.Email != nil {
		setParts = append(setParts, "email = :email")
		args["email"] = *personRequest.Email
	}

	// For password updates, you should handle hashing and verification outside this function.
	if personRequest.NewPassword != nil {
		setParts = append(setParts, "password = :password")
		args["password"] = *personRequest.NewPassword
	}

	// Add any other fields as necessary...

	if len(setParts) == 0 {
		return "", nil, fmt.Errorf("no fields to update")
	}

	// Always update the "updated_at" field.
	args["updated_at"] = time.Now().UTC().Format(time.RFC3339)
	setParts = append(setParts, "updated_at = :updated_at")

	sql := fmt.Sprintf("UPDATE persons SET %s WHERE id = :id", strings.Join(setParts, ", "))
	return sql, args, nil
}

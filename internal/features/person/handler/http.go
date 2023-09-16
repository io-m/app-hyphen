package person_http_adapter

import (
	"fmt"
	"net/http"

	person_objects "github.com/io-m/app-hyphen/internal/features/person/domain/objects"
	person_usecase "github.com/io-m/app-hyphen/internal/features/person/interface/usecase"
	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

type personHandler struct {
	personUsecase person_usecase.IPersonUsecase
	protector     tokens.IProtector
}

func NewPersonHandler(personUsecase person_usecase.IPersonUsecase, protector tokens.IProtector) *personHandler {
	return &personHandler{
		personUsecase: personUsecase,
		protector:     protector,
	}
}

func (ph *personHandler) GetById(w http.ResponseWriter, r *http.Request) {
	personId := helpers.GetUrlParam(r, "id")
	personResponse, err := ph.personUsecase.GetPersonWithId(r.Context(), personId)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not find person with id %s: %w", personId, err), http.StatusNotFound)
		return
	}
	helpers.SuccessResponse(w, personResponse, "Person found", http.StatusOK)
}
func (ph *personHandler) Update(w http.ResponseWriter, r *http.Request) {
	personId := helpers.GetUrlParam(r, "id")
	personRequest, err := helpers.DecodePayload[*person_objects.PersonRequestOptional](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}
	personResponse, err := ph.personUsecase.UpdatePersonWithId(r.Context(), personId, personRequest)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not update person with id %s: %w", personId, err), http.StatusInternalServerError)
		return
	}
	helpers.SuccessResponse(w, personResponse, "Person successfully updated", http.StatusNoContent)
}
func (ch *personHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	personId := helpers.GetUrlParam(r, "id")
	ok, err := ch.personUsecase.DeletePersonWithId(r.Context(), personId)
	if !ok || err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not delete person with id %s: %w", personId, err), http.StatusInternalServerError)
		return
	}
}

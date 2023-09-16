package auth_http_adapter

import (
	"fmt"
	"net/http"

	auth_usecase_interface "github.com/io-m/app-hyphen/internal/features/auth/interface/usecase"
	person_objects "github.com/io-m/app-hyphen/internal/features/person/domain/objects"
	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

type authHandler struct {
	authUsecase auth_usecase_interface.IAuthUsecase
	tokens      tokens.ITokens
	protector   tokens.IProtector
}

func NewAuthHandler(authUsecase auth_usecase_interface.IAuthUsecase, tokens tokens.ITokens, protector tokens.IProtector) *authHandler {
	return &authHandler{
		authUsecase: authUsecase,
		tokens:      tokens,
		protector:   protector,
	}
}

func (ah *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	personRequest, err := helpers.DecodePayload[*person_objects.PersonRequest](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}
	personResponse, err := ah.authUsecase.Register(r.Context(), personRequest)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not register: %w", err), http.StatusBadRequest)
		return
	}
	helpers.SuccessResponse(w, personResponse, "Person successfully registered", http.StatusCreated)
}

func (ah *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	c, err := helpers.DecodePayload[*person_objects.LoginPersonRequest](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}

	person, err := ah.authUsecase.Login(r.Context(), c)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not find person: %w", err), http.StatusNotFound)
		return
	}
	if err := helpers.CheckPassword(c.Password, person.Password); err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("wrong password: %w", err), http.StatusBadRequest)
		return
	}
	claims, _ := tokens.NewClaims(person.Id, constants.ACCESS_TOKEN_DURATION)

	accessToken, refreshToken, err := ah.protector.GenerateTokens(claims)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while generating tokens: %w", err), http.StatusInternalServerError)
		return
	}
	// Here we need to save refresh token in Redis
	if err := ah.tokens.SaveRefreshToken(r.Context(), person.Id, refreshToken); err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while saving refresh token: %w", err), http.StatusInternalServerError)
		return
	}
	w.Header().Add(constants.ACCESS_TOKEN_HEADER, accessToken)
	w.Header().Add(constants.REFRESH_TOKEN_HEADER, refreshToken)

	helpers.SuccessResponse(w, person_objects.MapPersonToPersonResponse(person), "Person successfully logged in")
}

func (ah *authHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}
func (ah *authHandler) OAuth(w http.ResponseWriter, r *http.Request) {

}

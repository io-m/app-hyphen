package types

import (
	"context"
	"net/http"
)

type ICrudHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
}

type ITokens interface {
	SaveRefreshToken(ctx context.Context, customerId, refreshToken string) error
	DeleteRefreshToken(ctx context.Context, customerId, refreshToken string) error
	RetrieveRefreshToken(ctx context.Context, customerId, refreshToken string) (string, error)
}

type ITokensIncoming interface {
	VerifyRefreshToken(ctx context.Context, customerId, refreshToken string) (bool, error)
}

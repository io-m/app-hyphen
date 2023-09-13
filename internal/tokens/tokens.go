package tokens

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/io-m/app-hyphen/pkg/constants"
)

type AuthorizationLevel string

const (
	CUSTOMER AuthorizationLevel = "CUSTOMER"
	PUBLIC   AuthorizationLevel = "PUBLIC"
	PROVIDER AuthorizationLevel = "PROVIDER"
	SUDO     AuthorizationLevel = "SUDO"
)

type Claims struct {
	ClaimID   uuid.UUID `json:"jti"`
	SubjectID uuid.UUID `json:"sub"`
	IssuedAt  time.Time `json:"iat"`
	ExpiredAt time.Time `json:"exp"`
	// Roles     []entities.AuthorizationLevel `json:"roles,omitempty"`
}

func NewClaims(subjectID uuid.UUID /*role entities.AuthorizationLevel,*/, duration time.Duration) (*Claims, error) {
	claimID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	claims := &Claims{
		ClaimID:   claimID,
		SubjectID: subjectID,
		IssuedAt:  time.Now().UTC(),
		ExpiredAt: time.Now().Add(constants.ACCESS_TOKEN_DURATION).UTC(),
		// Roles:     []entities.AuthorizationLevel{role},
	}
	return claims, nil
}

type ITokens interface {
	GenerateTokens(claims *Claims) (string, string, error)
	VerifyToken(token string) (*Claims, error)
	SaveRefreshToken(ctx context.Context, customerId uuid.UUID, refreshToken string) error
	DeleteRefreshToken(ctx context.Context, customerId uuid.UUID, refreshToken string) error
	RetrieveRefreshToken(ctx context.Context, customerId uuid.UUID, refreshToken string) (string, error)
	VerifyRefreshToken(ctx context.Context, customerId uuid.UUID, refreshToken string) (bool, error)
}

// Based on running environment we select authenticator
func NewAuthenticationTokens() ITokens {
	if os.Getenv(constants.RUNNING_ENV) == constants.DEVELOPMENT {
		return NewPasetoProtector()
	}
	return NewJWTProtector()
}

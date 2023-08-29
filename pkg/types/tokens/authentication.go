package tokens

import (
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
	SubjectID string    `json:"sub"`
	IssuedAt  time.Time `json:"iat"`
	ExpiredAt time.Time `json:"exp"`
	// Roles     []entities.AuthorizationLevel `json:"roles,omitempty"`
}

func NewClaims(subjectID string /*role entities.AuthorizationLevel,*/, duration time.Duration) (*Claims, error) {
	claimID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	claims := &Claims{
		ClaimID:   claimID,
		SubjectID: subjectID,
		IssuedAt:  time.Now().UTC(),
		ExpiredAt: time.Now().Add(duration).UTC(),
		// Roles:     []entities.AuthorizationLevel{role},
	}
	return claims, nil
}

type IAuthenticator interface {
	GenerateTokens(claims *Claims) (string, string, error)
	VerifyToken(token string) (*Claims, error)
}

// Based on running environment we select authenticator
func NewAuthenticator() IAuthenticator {
	if os.Getenv(constants.RUNNING_ENV) == constants.DEVELOPMENT {
		return NewPasetoProtector()
	}
	return NewJWTProtector()
}

package tokens

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/io-m/app-hyphen/pkg/constants"
)

type IProtector interface {
	GenerateTokens(claims *Claims) (string, string, error)
	VerifyToken(token string) (*Claims, error)
	VerifyRefreshToken(ctx context.Context, customerId uuid.UUID, refreshToken string) (bool, error)
}

// Based on running environment we select protector
func NewProtector() IProtector {
	if os.Getenv(constants.RUNNING_ENV) == constants.DEVELOPMENT {
		return NewPasetoProtector()
	}
	return NewJWTProtector()
}

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

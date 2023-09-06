package tokens

import (
	"context"

	"github.com/io-m/app-hyphen/pkg/constants"
)

type jwtAuthenticator struct {
	accessTokenSecretKey  []byte
	refreshTokenSecretKey []byte
}

func NewJWTProtector() ITokens {
	return &jwtAuthenticator{
		accessTokenSecretKey:  []byte(constants.ACCESS_TOKEN_SECRET_KEY),
		refreshTokenSecretKey: []byte(constants.REFRESH_TOKEN_SECRET_KEY),
	}
}

// TODO: Implement NewJWTProtector to be IAuthenticator
func (protector *jwtAuthenticator) GenerateTokens(claims *Claims) (string, string, error) {

	return "", "", nil
}

// TODO: Implement NewJWTProtector to be IAuthenticator
func (protector *jwtAuthenticator) VerifyToken(stringifiedToken string) (*Claims, error) {

	return nil, nil
}

func (protector *jwtAuthenticator) SaveRefreshToken(ctx context.Context, customerId, refreshToken string) error {
	return nil
}
func (protector *jwtAuthenticator) DeleteRefreshToken(ctx context.Context, customerId, refreshToken string) error {
	return nil
}
func (protector *jwtAuthenticator) RetrieveRefreshToken(ctx context.Context, customerId, refreshToken string) (string, error) {
	return "", nil
}
func (protector *jwtAuthenticator) VerifyRefreshToken(ctx context.Context, customerId, refreshToken string) (bool, error) {
	return false, nil
}

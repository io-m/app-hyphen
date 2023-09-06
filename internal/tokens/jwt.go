package tokens

import (
	"github.com/io-m/app-hyphen/pkg/constants"
)

type jwtAuthenticator struct {
	accessTokenSecretKey  []byte
	refreshTokenSecretKey []byte
}

func NewJWTProtector() IAuthenticator {
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

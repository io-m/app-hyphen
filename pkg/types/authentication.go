package types

import (
	"os"

	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/tokens"
)

type IAuthenticator interface {
	GenerateTokens(claims *tokens.Claims) (string, string, error)
	VerifyToken(token string) (*tokens.Claims, error)
}

// Based on running environment we select authenticator
func NewAuthenticator() IAuthenticator {
	if os.Getenv(constants.RUNNING_ENV) == constants.DEVELOPMENT {
		return tokens.NewPasetoProtector()
	}
	// TODO: Implement NewJWTProtector to be IAuthenticator
	// return NewJWTProtector()
	return nil
}

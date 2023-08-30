package customer_logic

import (
	"github.com/io-m/app-hyphen/pkg/types/tokens"
)

func (cl *customerLogic) GenerateTokens(claims *tokens.Claims) (string, string, error) {
	return cl.authenticator.GenerateTokens(claims)
}

func (cl *customerLogic) VerifyToken(token string) (*tokens.Claims, error) {
	return cl.authenticator.VerifyToken(token)
}

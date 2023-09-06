package tokens

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"aidanwoods.dev/go-paseto"
	"github.com/io-m/app-hyphen/pkg/constants"
)

type pasetoProtector struct {
	token                 paseto.Token
	parser                paseto.Parser
	accessTokenSecretKey  []byte
	refreshTokenSecretKey []byte
}

func NewPasetoProtector() *pasetoProtector {
	return &pasetoProtector{
		token:                 paseto.NewToken(),
		parser:                paseto.NewParser(),
		accessTokenSecretKey:  []byte(os.Getenv(constants.ACCESS_TOKEN_SECRET_KEY)),
		refreshTokenSecretKey: []byte(os.Getenv(constants.REFRESH_TOKEN_SECRET_KEY)),
	}
}

func (protector *pasetoProtector) VerifyToken(stringifiedToken string) (*Claims, error) {
	symmetricAccessTokenKey, err := paseto.V4SymmetricKeyFromBytes(protector.accessTokenSecretKey)
	if err != nil {
		return nil, fmt.Errorf("paseto could not generate V4Symmetric key for access token: %w", err)
	}
	token, err := protector.parser.ParseV4Local(symmetricAccessTokenKey, stringifiedToken, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot extract paseto token %w", err)
	}
	claimsBytes := token.ClaimsJSON()

	var claims *Claims
	if err := json.Unmarshal(claimsBytes, &claims); err != nil {
		return nil, err
	}
	return claims, nil
}

func (protector *pasetoProtector) GenerateTokens(claims *Claims) (string, string, error) {
	protector.token.SetExpiration(claims.ExpiredAt)
	protector.token.SetIssuedAt(claims.IssuedAt)
	protector.token.SetSubject(claims.SubjectID)
	protector.token.SetJti(claims.ClaimID.String())
	// protector.token.Set("roles", claims.Roles)
	symmetricAccessTokenKey, err := paseto.V4SymmetricKeyFromBytes(protector.accessTokenSecretKey)
	if err != nil {
		return "", "", fmt.Errorf("paseto could not generate V4Symmetric key for access token: %w", err)
	}
	symmetricRefreshTokenKey, err := paseto.V4SymmetricKeyFromBytes(protector.refreshTokenSecretKey)
	if err != nil {
		return "", "", fmt.Errorf("paseto could not generate V4Symmetric key for refresh token: %w", err)
	}
	encryptedAccessTokenKey := protector.token.V4Encrypt(symmetricAccessTokenKey, nil)
	encryptedRefreshTokenKey := protector.token.V4Encrypt(symmetricRefreshTokenKey, nil)
	return encryptedAccessTokenKey, encryptedRefreshTokenKey, nil
}

func (protector *pasetoProtector) SaveRefreshToken(ctx context.Context, customerId, refreshToken string) error {
	return nil
}
func (protector *pasetoProtector) DeleteRefreshToken(ctx context.Context, customerId, refreshToken string) error {
	return nil
}
func (protector *pasetoProtector) RetrieveRefreshToken(ctx context.Context, customerId, refreshToken string) (string, error) {
	return "", nil
}
func (protector *pasetoProtector) VerifyRefreshToken(ctx context.Context, customerId, refreshToken string) (bool, error) {
	return false, nil
}

package helpers

import (
	"context"

	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/io-m/app-hyphen/pkg/constants"
)

// IsUserAuthorized compares id extracted from url params and subject id stored in claims
func IsUserAuthorized(ctx context.Context, id string) bool {
	claims, ok := ctx.Value(constants.CLAIMS).(*tokens.Claims)
	if !ok {
		return false
	}
	if claims.SubjectID != id {
		return false
	}
	return true
}

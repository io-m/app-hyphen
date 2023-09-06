package types

import (
	"github.com/go-chi/chi/v5"
	"github.com/io-m/app-hyphen/internal/tokens"
)

type AppConfig struct {
	Mux           chi.Router
	Authenticator tokens.ITokens
}

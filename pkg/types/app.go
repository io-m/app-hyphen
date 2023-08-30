package types

import (
	"github.com/go-chi/chi/v5"
	"github.com/io-m/app-hyphen/pkg/types/tokens"
)

type AppConfig struct {
	Mux           chi.Router
	Authenticator tokens.IAuthenticator
}

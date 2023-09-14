package auth_routes

import (
	"github.com/go-chi/chi/v5"
	auth_http_adapter "github.com/io-m/app-hyphen/internal/auth/handler"
	auth_usecase "github.com/io-m/app-hyphen/internal/auth/usecase"
	"github.com/io-m/app-hyphen/pkg/types"
)

func SetAndRunAuthRoutes(config *types.AppConfig) {
	authUsecase := auth_usecase.NewAuthUsecase(config.Repositories.AddressRepository, config.Repositories.CustomerRepository)
	authHandler := auth_http_adapter.NewAuthHandler(authUsecase, config.Tokens, config.Protector)

	/* AUTH ROUTES */
	config.Router.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/oauth", authHandler.OAuth)
		r.Post("/refresh-tokens", authHandler.RefreshToken)

	})
}

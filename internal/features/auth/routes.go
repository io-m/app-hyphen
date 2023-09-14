package auth_routes

import (
	"github.com/go-chi/chi/v5"
	auth_http_adapter "github.com/io-m/app-hyphen/internal/features/auth/handler"
	auth_usecase "github.com/io-m/app-hyphen/internal/features/auth/usecase"
	"github.com/io-m/app-hyphen/internal/shared"
)

func SetAndRunAuthRoutes(config *shared.AppConfig) {
	authUsecase := auth_usecase.NewAuthUsecase(config.GetAddressRepository(), config.GetCustomerRepository())
	authHandler := auth_http_adapter.NewAuthHandler(authUsecase, config.GetTokens(), config.GetProtector())

	/* AUTH ROUTES */
	config.GetRouter().Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/oauth", authHandler.OAuth)
		r.Post("/refresh-tokens", authHandler.RefreshToken)

	})
}

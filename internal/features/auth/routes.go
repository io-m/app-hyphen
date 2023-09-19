package auth_routes

import (
	"github.com/go-chi/chi/v5"
	auth_http_adapter "github.com/io-m/app-hyphen/internal/features/auth/handler"
	auth_usecase "github.com/io-m/app-hyphen/internal/features/auth/usecase"
	"github.com/io-m/app-hyphen/internal/shared"
)

func SetAndRunAuthRoutes(config *shared.AppConfig) {
	authUsecase := auth_usecase.NewAuthUsecase(config.GetPersonRepository())
	authHandler := auth_http_adapter.NewAuthHandler(authUsecase, config.GetTokens(), config.GetProtector())

	/* AUTH ROUTES */
	config.GetRouter().Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/refresh-tokens", authHandler.RefreshToken)
		r.Put("/password-reset", authHandler.RefreshToken) // TODO: Implement real handler
		r.Route("/oauth", func(r chi.Router) {
			r.Get("/{provider}/login", authHandler.OAuth)    // TODO: Implement real handler
			r.Get("/{provider}/callback", authHandler.OAuth) // TODO: Implement real handler
		})

	})
}

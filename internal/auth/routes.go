package auth_routes

import (
	"github.com/go-chi/chi/v5"
	address_db_repository "github.com/io-m/app-hyphen/internal/address/repository"
	auth_http_adapter "github.com/io-m/app-hyphen/internal/auth/handler"
	auth_usecase "github.com/io-m/app-hyphen/internal/auth/usecase"
	customer_db_adapter "github.com/io-m/app-hyphen/internal/customer/repository"
	"github.com/io-m/app-hyphen/pkg/types"
)

func SetAndRunCustomerRoutes(config *types.AppConfig) {
	addressRepository := address_db_repository.NewAddressRepository(config.Postgres, nil)

	customerRepository := customer_db_adapter.NewCustomerRepository(config.Postgres, config.RedisClient)
	authUsecase := auth_usecase.NewAuthUsecase(addressRepository, customerRepository)
	authHandler := auth_http_adapter.NewAuthHandler(authUsecase, config.Tokens, config.Protector)

	/* CUSTOMER ROUTES */
	config.Router.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/oauth", authHandler.OAuth)
		r.Post("/refresh-tokens", authHandler.RefreshToken)

	})
}

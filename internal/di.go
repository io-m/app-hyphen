package di

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	auth_routes "github.com/io-m/app-hyphen/internal/auth"
	customer_common "github.com/io-m/app-hyphen/internal/customer"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/types"
)

func ConfigureRoutes(config *types.AppConfig) {
	// authenticator := tokens.NewAuthenticationTokens()
	// mux := chi.NewRouter()
	config.Mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	config.Mux.Use(middleware.Heartbeat("/ping"))
	config.Mux.Use(middleware.Logger)
	config.Mux.Use(middleware.Recoverer)
	config.Mux.Route(constants.BASE_ROUTE, func(r chi.Router) {
		config.Router = r
		/* ROUTES COME HERE*/
		auth_routes.SetAndRunAuthRoutes(config)
		customer_common.SetAndRunCustomerRoutes(config)
	})
}

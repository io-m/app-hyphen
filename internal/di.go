package di

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	auth_routes "github.com/io-m/app-hyphen/internal/features/auth"
	customer_routes "github.com/io-m/app-hyphen/internal/features/customer"
	"github.com/io-m/app-hyphen/internal/shared"
	"github.com/io-m/app-hyphen/pkg/constants"
)

func ConfigureRoutes(config *shared.AppConfig) {
	config.GetMux().Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	config.GetMux().Use(middleware.Heartbeat("/ping"))
	config.GetMux().Use(middleware.Logger)
	config.GetMux().Use(middleware.Recoverer)
	config.GetMux().Route(constants.BASE_ROUTE, func(r chi.Router) {
		config.SetRouter(r)
		/* ROUTES COME HERE*/
		auth_routes.SetAndRunAuthRoutes(config)
		customer_routes.SetAndRunCustomerRoutes(config)
	})
}

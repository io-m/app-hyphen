package di

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	customer_common "github.com/io-m/app-hyphen/internal/customer"
	"github.com/io-m/app-hyphen/pkg/helpers"
	"github.com/io-m/app-hyphen/pkg/types"
	"github.com/io-m/app-hyphen/pkg/types/tokens"
)

func SetAndRun() *chi.Mux {
	arangoDriver, err := helpers.CreateArangoConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	mux := chi.NewRouter()
	config := &types.AppConfig{
		Mux:           mux,
		Authenticator: tokens.NewAuthenticator(),
	}
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	customer_common.SetAndRunCustomerRoutes(config, arangoDriver)

	return mux
}

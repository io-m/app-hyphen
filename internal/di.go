package di

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	customer_common "github.com/io-m/app-hyphen/internal/customer"
	"github.com/io-m/app-hyphen/pkg/helpers"
)

func SetAndRun() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	// mux.Use(middlewares.JSONMiddleware)
	arangoDriver, err := helpers.CreateArangoConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	customer_common.SetAndRunCustomerRoutes(mux, arangoDriver)

	return mux
}

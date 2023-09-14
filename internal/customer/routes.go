package customer_routes

import (
	"github.com/go-chi/chi/v5"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/customer/handler"
	customer_usecase_adapter "github.com/io-m/app-hyphen/internal/customer/usecase"
	"github.com/io-m/app-hyphen/pkg/middlewares"
	"github.com/io-m/app-hyphen/pkg/types"
)

func SetAndRunCustomerRoutes(config *types.AppConfig) {
	customerUsecase := customer_usecase_adapter.NewCustomerUsecase(config.Repositories.AddressRepository, config.Repositories.CustomerRepository, config.Tokens)
	customerHandler := customer_http_adapter.NewCustomerHandler(customerUsecase, config.Protector)

	/* CUSTOMER ROUTES */
	config.Router.Route("/customers", func(r chi.Router) {
		r.Get("/{id}", customerHandler.GetById)
		/// Authentication required
		r.Route("/", func(r chi.Router) {
			r.Use(middlewares.MustAuthenticate(config.Protector))
			r.Put("/{id}", customerHandler.Update)
			r.Delete("/{id}", customerHandler.DeleteById)
		})
	})
}

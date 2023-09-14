package customer_routes

import (
	"github.com/go-chi/chi/v5"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/features/customer/handler"
	customer_usecase_adapter "github.com/io-m/app-hyphen/internal/features/customer/usecase"
	"github.com/io-m/app-hyphen/internal/shared"
	"github.com/io-m/app-hyphen/pkg/middlewares"
)

func SetAndRunCustomerRoutes(config *shared.AppConfig) {
	customerUsecase := customer_usecase_adapter.NewCustomerUsecase(config.GetAddressRepository(), config.GetCustomerRepository(), config.GetTokens())
	customerHandler := customer_http_adapter.NewCustomerHandler(customerUsecase, config.GetProtector())

	/* CUSTOMER ROUTES */
	config.GetRouter().Route("/customers", func(r chi.Router) {
		r.Get("/{id}", customerHandler.GetById)
		/// Authentication required
		r.Route("/", func(r chi.Router) {
			r.Use(middlewares.MustAuthenticate(config.GetProtector()))
			r.Put("/{id}", customerHandler.Update)
			r.Delete("/{id}", customerHandler.DeleteById)
		})
	})
}

package customer_routes

import (
	"github.com/go-chi/chi/v5"
	address_db_repository "github.com/io-m/app-hyphen/internal/address/adapters/database"
	customer_db_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/database"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/http"
	customer_usecase_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/usecase"
	"github.com/io-m/app-hyphen/pkg/middlewares"
	"github.com/io-m/app-hyphen/pkg/types"
)

func SetAndRunCustomerRoutes(config *types.AppConfig) {
	addressRepository := address_db_repository.NewAddressRepository(config.Postgres, nil)

	customerRepository := customer_db_adapter.NewCustomerRepository(config.Postgres, config.RedisClient)
	customerUsecase := customer_usecase_adapter.NewCustomerUsecase(addressRepository, customerRepository, config.Authenticator)
	customerHandler := customer_http_adapter.NewCustomerRESTHandler(customerUsecase, config.Authenticator)

	/* CUSTOMER ROUTES */
	config.Router.Route("/customers", func(r chi.Router) {
		r.Get("/{id}", customerHandler.GetById)
		r.Post("/register", customerHandler.Create)
		r.Post("/login", customerHandler.Login)

		/// Authentication required
		r.Route("/", func(r chi.Router) {
			r.Use(middlewares.MustAuthenticate(config.Authenticator))
			r.Put("/{id}", customerHandler.Update)
			r.Delete("/{id}", customerHandler.DeleteById)
		})
	})
}

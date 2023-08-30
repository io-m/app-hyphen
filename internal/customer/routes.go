package customer_common

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/http"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/middlewares"
	"github.com/io-m/app-hyphen/pkg/types"
)

func handleCustomerRoutes(config *types.AppConfig, handler *customer_http_adapter.CustomerRESTHandler) {
	config.Mux.Route(fmt.Sprintf("%s/customers", constants.BASE_ROUTE), func(r chi.Router) {
		r.Get("/", handler.GetAllCustomers)
		// r.Get("/{id}", handler.GetCustomerById)
		r.Post("/register", handler.CreateCustomer)
		// r.Post("/login", hyphen.Login)

		/// Authentication required
		r.Route("/", func(r chi.Router) {
			r.Use(middlewares.MustAuthenticate(config.Authenticator))
			r.Get("/{id}", handler.GetCustomerById)
			r.Put("/{customer_id}", handler.UpdateCustomer)
			r.Delete("/{customer_id}", handler.DeleteCustomerById)
		})
	})

}

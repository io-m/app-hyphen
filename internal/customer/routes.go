package customer_common

import (
	"github.com/go-chi/chi/v5"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/http"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/types"
)

func handleCustomerRoutes(config *types.AppConfig, handler *customer_http_adapter.CustomerRESTHandler) {
	config.Mux.Route(constants.BASE_ROUTE, func(r chi.Router) {
		r.Route("/customers", func(r chi.Router) {
			// r.Use(middlewares.SudoOnly)    ---> This way we can set some router level middleware
			r.Get("/", handler.GetAllCustomers)
			r.Get("/{id}", handler.GetCustomerById)
			r.Post("/register", handler.CreateCustomer)
			// r.Post("/login", hyphen.Login)
			/// Authentication required
			// TODO: implement types.IAuthenticator
			// r.Use(middlewares.MustAuthenticate(config.Authenticator))
			r.Put("/api/book/{book_id}", handler.UpdateCustomer)
			r.Delete("/api/book/{book_id}", handler.DeleteCustomerById)
		})
	})
}

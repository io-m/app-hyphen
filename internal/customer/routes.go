package customer_common

import (
	"github.com/arangodb/go-driver"
	"github.com/go-chi/chi/v5"
	customer_arango_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/database"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/http"
	customer_logic "github.com/io-m/app-hyphen/internal/customer/logic"
	"github.com/io-m/app-hyphen/pkg/middlewares"
	"github.com/io-m/app-hyphen/pkg/types"
)

func SetAndRunCustomerRoutes(config *types.AppConfig, arangoDriver driver.Database) {
	// TODO: Set Redis adapter and pass it to NewCustomerDB
	dbAdapter := customer_arango_adapter.NewCustomerDB(arangoDriver, struct{}{})
	customerLogic := customer_logic.NewCustomerLogic(dbAdapter)
	customerHandler := customer_http_adapter.NewCustomerRESTHandler(customerLogic)

	/* CUSTOMER ROUTES */
	config.Mux.Route("/customers", func(r chi.Router) {
		r.Get("/", customerHandler.GetAllCustomers)
		r.Get("/{id}", customerHandler.GetCustomerById)
		r.Post("/register", customerHandler.CreateCustomer)
		// r.Post("/login", hyphen.Login)

		/// Authentication required
		r.Route("/", func(r chi.Router) {
			r.Use(middlewares.MustAuthenticate(config.Authenticator))
			r.Put("/{customer_id}", customerHandler.UpdateCustomer)
			r.Delete("/{customer_id}", customerHandler.DeleteCustomerById)
		})
	})
}

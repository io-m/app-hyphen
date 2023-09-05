package customer_routes

import (
	"github.com/arangodb/go-driver"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	customer_db_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/database"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/http"
	"github.com/io-m/app-hyphen/pkg/middlewares"
	"github.com/io-m/app-hyphen/pkg/types"
)

func SetAndRunCustomerRoutes(config *types.AppConfig, arangoDriver driver.Database, redisClient *redis.Client) {
	customerRepository := customer_db_adapter.NewCustomerRepository(arangoDriver, redisClient)
	customerHandler := customer_http_adapter.NewCustomerRESTHandler(customerRepository, config.Authenticator)

	/* CUSTOMER ROUTES */
	config.Mux.Route("/customers", func(r chi.Router) {
		// r.Get("/", customerHandler.GetAllCustomers)
		r.Get("/{id}", customerHandler.GetById)
		r.Post("/register", customerHandler.Create)
		r.Post("/login", customerHandler.LoginCustomer)

		/// Authentication required
		r.Route("/", func(r chi.Router) {
			r.Use(middlewares.MustAuthenticate(config.Authenticator))
			r.Put("/{id}", customerHandler.Update)
			r.Delete("/{id}", customerHandler.DeleteById)
		})
	})
}

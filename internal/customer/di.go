package customer_common

import (
	"github.com/arangodb/go-driver"
	"github.com/go-chi/chi/v5"
	customer_arango_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/arango"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/http"
	customer_logic "github.com/io-m/app-hyphen/internal/customer/logic"
)

func SetAndRunCustomerRoutes(mux *chi.Mux, arangoDriver driver.Database) {
	arangoAdapter := customer_arango_adapter.NewArangoCustomerDB(arangoDriver)
	customerLogic := customer_logic.NewCustomerLogic(arangoAdapter)
	customerHandler := customer_http_adapter.NewCustomerRESTHandler(customerLogic)

	HandleCustomerRoutes(mux, customerHandler)

}

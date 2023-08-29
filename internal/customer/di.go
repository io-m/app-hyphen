package customer_common

import (
	"github.com/arangodb/go-driver"
	customer_arango_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/arango"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/http"
	customer_logic "github.com/io-m/app-hyphen/internal/customer/logic"
	"github.com/io-m/app-hyphen/pkg/types"
)

func SetAndRunCustomerRoutes(config *types.AppConfig, arangoDriver driver.Database) {
	arangoAdapter := customer_arango_adapter.NewArangoCustomerDB(arangoDriver)
	customerLogic := customer_logic.NewCustomerLogic(arangoAdapter)
	customerHandler := customer_http_adapter.NewCustomerRESTHandler(customerLogic)

	handleCustomerRoutes(config, customerHandler)
}

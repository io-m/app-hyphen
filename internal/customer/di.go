package customer_common

import (
	"github.com/arangodb/go-driver"
	customer_arango_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/database"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/http"
	customer_logic "github.com/io-m/app-hyphen/internal/customer/logic"
	"github.com/io-m/app-hyphen/pkg/types"
)

func SetAndRunCustomerRoutes(config *types.AppConfig, arangoDriver driver.Database) {
	dbAdapter := customer_arango_adapter.NewCustomerDB(arangoDriver, struct{}{})
	customerLogic := customer_logic.NewCustomerLogic(dbAdapter)
	customerHandler := customer_http_adapter.NewCustomerRESTHandler(customerLogic)

	handleCustomerRoutes(config, customerHandler)
}

package queries

import (
	"log"
	"strings"
	"time"

	customer_objects "github.com/io-m/app-hyphen/internal/customer/domain/objects"
)

const CREATE_CUSTOMER_QUERY = "INSERT @customer INTO customers"
const GET_CUSTOMER_BY_ID_QUERY = `
	FOR customer IN customers
	FILTER customer.id == @customerId
	RETURN customer
`

func BuildUpdateQueryAndVars(req *customer_objects.CustomerRequestOptional, customerId string) (string, map[string]interface{}) {
	log.Println("ID ==> ", customerId)
	baseQuery := `
	LET customerData = {`

	bindVars := make(map[string]interface{})

	// Always add the CustomerID as that is required for filtering
	bindVars["CustomerID"] = customerId
	bindVars["UpdatedAt"] = time.Now().Format(time.RFC3339)
	baseQuery += `
		"updated_at": @UpdatedAt,`
	if req.FirstName != nil {
		baseQuery += `
		"first_name": @FirstName,`
		bindVars["FirstName"] = *req.FirstName
	}

	if req.LastName != nil {
		baseQuery += `
		"last_name": @LastName,`
		bindVars["LastName"] = *req.LastName
	}

	if req.Email != nil {
		baseQuery += `
		"email": @Email,`
		bindVars["Email"] = *req.Email
	}

	if req.NewPassword != nil {
		baseQuery += `
		"password": @NewPassword,`

		bindVars["NewPassword"] = req.NewPassword
	}

	if req.Address != nil {
		baseQuery += `
		"address": @Address,`
		bindVars["Address"] = *req.Address
	}

	if req.Role != nil {
		baseQuery += `
		"role": @Role,`
		bindVars["Role"] = *req.Role
	}

	// Remove trailing comma and close the opening brace for customerData
	baseQuery = strings.TrimSuffix(baseQuery, ",")
	baseQuery += `
	}
	FOR customer IN customers
	FILTER customer.id == @CustomerID
	UPDATE customer WITH customerData IN customers
	RETURN NEW
	`

	return baseQuery, bindVars
}

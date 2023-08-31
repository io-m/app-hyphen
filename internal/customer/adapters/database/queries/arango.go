package queries

const CREATE_CUSTOMER_QUERY = "INSERT @customer INTO customers"
const GET_CUSTOMER_BY_ID_QUERY = `
	FOR customer IN customers
	FILTER customer.id == @customerId
	RETURN customer
`

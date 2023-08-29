package customer_common

import (
	"github.com/go-chi/chi/v5"
	customer_http_adapter "github.com/io-m/app-hyphen/internal/customer/adapters/http"
)

func HandleCustomerRoutes(mux *chi.Mux, handler *customer_http_adapter.CustomerRESTHandler) {
	mux.Post("/api/book", handler.CreateCustomer)
	mux.Get("/api/book/{book_id}", handler.GetCustomerById)
	mux.Get("/api/book/books", handler.GetAllCustomers)
	mux.Put("/api/book/{book_id}", handler.UpdateCustomer)
	mux.Delete("/api/book/{book_id}", handler.DeleteCustomerById)
}

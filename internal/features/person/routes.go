package person_routes

import (
	"github.com/go-chi/chi/v5"
	person_http_adapter "github.com/io-m/app-hyphen/internal/features/person/handler"
	person_usecase_adapter "github.com/io-m/app-hyphen/internal/features/person/usecase"
	"github.com/io-m/app-hyphen/internal/shared"
	"github.com/io-m/app-hyphen/internal/shared/middlewares"
)

func SetAndRunPersonRoutes(config *shared.AppConfig) {
	personUsecase := person_usecase_adapter.NewPersonUsecase(config.GetPersonRepository(), config.GetTokens())
	personHandler := person_http_adapter.NewPersonHandler(personUsecase, config.GetProtector())

	/* PERSON ROUTES */
	config.GetRouter().Route("/persons", func(r chi.Router) {
		r.Get("/{id}", personHandler.GetById)
		/// Authentication required
		r.Route("/", func(r chi.Router) {
			r.Use(middlewares.MustAuthenticate(config.GetProtector()))
			r.Put("/{id}", personHandler.Update)
			r.Delete("/{id}", personHandler.DeleteById)
		})
	})
}

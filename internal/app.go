package di

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	customer_common "github.com/io-m/app-hyphen/internal/customer"
	hyphen_arango "github.com/io-m/app-hyphen/pkg/arango"
	"github.com/io-m/app-hyphen/pkg/constants"
	hyphen_redis "github.com/io-m/app-hyphen/pkg/redis"
	"github.com/io-m/app-hyphen/pkg/types"
	"github.com/io-m/app-hyphen/pkg/types/tokens"
)

func SetAndRun() *chi.Mux {
	arangoDriver, err := hyphen_arango.CreateArangoConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	redisClient, err := hyphen_redis.CreateRedisConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Deferring Redis conn closing
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	authenticator := tokens.NewAuthenticator()
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Route(constants.BASE_ROUTE, func(r chi.Router) {
		config := &types.AppConfig{
			Mux:           r,
			Authenticator: authenticator,
		}
		/* ROUTES COME HERE*/
		customer_common.SetAndRunCustomerRoutes(config, arangoDriver, redisClient)
	})

	return mux
}

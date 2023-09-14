package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	di "github.com/io-m/app-hyphen/internal"
	address_repository "github.com/io-m/app-hyphen/internal/address/repository"
	customer_repository "github.com/io-m/app-hyphen/internal/customer/repository"
	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/helpers"
	"github.com/io-m/app-hyphen/pkg/postgres"
	hyphen_redis "github.com/io-m/app-hyphen/pkg/redis"
	"github.com/io-m/app-hyphen/pkg/repositories"
	"github.com/io-m/app-hyphen/pkg/types"
	"golang.org/x/exp/slog"
)

func main() {
	helpers.LoadEnv(constants.PROD_CONFIG_FILE)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	postgresConnection, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Deferring Postgres conn closing
	defer func() {
		if err := postgresConnection.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

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

	protector := tokens.NewProtector()
	mux := chi.NewRouter()
	repositories := &repositories.AppRepositories{
		AddressRepository:  address_repository.NewAddressRepository(postgresConnection, redisClient),
		CustomerRepository: customer_repository.NewCustomerRepository(postgresConnection, redisClient),
	}
	// Global singleton
	config := &types.AppConfig{
		Mux:          mux,
		Protector:    protector,
		Tokens:       tokens.NewTokens(redisClient),
		Postgres:     postgresConnection,
		RedisClient:  redisClient,
		Repositories: repositories,
	}

	port := os.Getenv(constants.APP_PORT)
	log.Printf("listening on port: %s............\n", port)
	di.ConfigureRoutes(config)
	go func() {
		runServer(port, mux)
	}()

	// Handling graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("Shutting down gracefully...")
}

// Function for running server
func runServer(port string, mux *chi.Mux) {
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("Server is down!")
	}
}

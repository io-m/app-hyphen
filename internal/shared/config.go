package shared

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	address_repository_interface "github.com/io-m/app-hyphen/internal/features/address/interface/repository"
	address_repository "github.com/io-m/app-hyphen/internal/features/address/repository"
	customer_repository_interface "github.com/io-m/app-hyphen/internal/features/customer/interface/repository"
	customer_repository "github.com/io-m/app-hyphen/internal/features/customer/repository"
	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/jmoiron/sqlx"
)

type AppConfig struct {
	mux                *chi.Mux
	router             chi.Router
	protector          tokens.IProtector
	tokens             tokens.ITokens
	postgres           *sqlx.DB
	redisClient        *redis.Client
	addressRepository  address_repository_interface.IAddressRepository
	customerRepository customer_repository_interface.ICustomerRepository
}

func NewAppConfig(pg *sqlx.DB, redis *redis.Client) *AppConfig {

	return &AppConfig{
		mux:                chi.NewRouter(),
		protector:          tokens.NewProtector(),
		tokens:             tokens.NewTokens(redis),
		postgres:           pg,
		redisClient:        redis,
		addressRepository:  address_repository.NewAddressRepository(pg, redis),
		customerRepository: customer_repository.NewCustomerRepository(pg, redis),
	}
}

func (ac *AppConfig) GetRouter() chi.Router {
	return ac.router
}
func (ac *AppConfig) GetMux() *chi.Mux {
	return ac.mux
}
func (ac *AppConfig) GetProtector() tokens.IProtector {
	return ac.protector
}
func (ac *AppConfig) GetTokens() tokens.ITokens {
	return ac.tokens
}
func (ac *AppConfig) GetPostgres() *sqlx.DB {
	return ac.postgres
}
func (ac *AppConfig) GetRedis() *redis.Client {
	return ac.redisClient
}
func (ac *AppConfig) GetAddressRepository() address_repository_interface.IAddressRepository {
	return ac.addressRepository
}
func (ac *AppConfig) GetCustomerRepository() customer_repository_interface.ICustomerRepository {
	return ac.customerRepository
}
func (ac *AppConfig) SetRouter(router chi.Router) {
	ac.router = router
}

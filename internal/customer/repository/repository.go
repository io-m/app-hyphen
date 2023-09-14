package customer_repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type customerRepository struct {
	postgres *sqlx.DB
	redis    *redis.Client
}

func NewCustomerRepository(postgres *sqlx.DB, redisClient *redis.Client) *customerRepository {
	return &customerRepository{
		postgres: postgres,
		redis:    redisClient,
	}
}

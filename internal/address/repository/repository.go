package address_db_repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type addressRepository struct {
	postgres *sqlx.DB
	redis    *redis.Client
}

func NewAddressRepository(postgres *sqlx.DB, redisClient *redis.Client) *addressRepository {
	return &addressRepository{
		postgres: postgres,
		redis:    redisClient,
	}
}

package person_repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type personRepository struct {
	postgres *sqlx.DB
	redis    *redis.Client
}

func NewPersonRepository(postgres *sqlx.DB, redisClient *redis.Client) *personRepository {
	return &personRepository{
		postgres: postgres,
		redis:    redisClient,
	}
}

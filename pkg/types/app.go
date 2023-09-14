package types

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/io-m/app-hyphen/internal/tokens"
	"github.com/jmoiron/sqlx"
)

type AppConfig struct {
	Mux         *chi.Mux
	Router      chi.Router
	Protector   tokens.IProtector
	Tokens      tokens.ITokens
	Postgres    *sqlx.DB
	RedisClient *redis.Client
}

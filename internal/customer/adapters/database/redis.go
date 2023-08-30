package customer_db_adapter

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/io-m/app-hyphen/pkg/constants"
)

func (db *customerOutgoing) SaveRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	err := db.redis.Set(ctx, constants.REFRESH_TOKEN_KEY, refreshToken, constants.REFRESH_TOKEN_DURATION)
	if err != nil {
		return "", err.Err()
	}

	return "success", nil
}

func (db *customerOutgoing) RetrieveAndVerifyRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	rt, err := db.redis.Get(ctx, constants.REFRESH_TOKEN_KEY).Result()
	if err != redis.Nil {
		return "", fmt.Errorf("%s key does not exist: %w", constants.REFRESH_TOKEN_KEY, err)
	} else if err != nil {
		return "", err
	}

	return rt, nil
}

package customer_repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/io-m/app-hyphen/pkg/constants"
)

func (db *customerRepository) SaveRefreshToken(ctx context.Context, customerId uuid.UUID, refreshToken string) error {
	err := db.redis.HSet(ctx, customerId.String(), "refreshToken", refreshToken)
	// err := db.redis.Set(ctx, constants.REFRESH_TOKEN_KEY, refreshToken, constants.REFRESH_TOKEN_DURATION)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (db *customerRepository) RetrieveRefreshToken(ctx context.Context, customerId uuid.UUID, refreshToken string) (string, error) {
	rt, err := db.redis.Get(ctx, constants.REFRESH_TOKEN_KEY).Result()
	if err != redis.Nil {
		return "", fmt.Errorf("%s key does not exist: %w", constants.REFRESH_TOKEN_KEY, err)
	} else if err != nil {
		return "", err
	}
	return rt, nil
}

func (db *customerRepository) DeleteRefreshToken(ctx context.Context, customerId uuid.UUID, refreshToken string) error {
	return nil
}

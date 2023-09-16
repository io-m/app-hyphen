package person_repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/io-m/app-hyphen/pkg/constants"
)

func (db *personRepository) SaveRefreshToken(ctx context.Context, personId uuid.UUID, refreshToken string) error {
	err := db.redis.HSet(ctx, personId.String(), "refreshToken", refreshToken)
	// err := db.redis.Set(ctx, constants.REFRESH_TOKEN_KEY, refreshToken, constants.REFRESH_TOKEN_DURATION)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (db *personRepository) RetrieveRefreshToken(ctx context.Context, personId uuid.UUID, refreshToken string) (string, error) {
	rt, err := db.redis.Get(ctx, constants.REFRESH_TOKEN_KEY).Result()
	if err != redis.Nil {
		return "", fmt.Errorf("%s refresh key does not exist: %w", constants.REFRESH_TOKEN_KEY, err)
	} else if err != nil {
		return "", err
	}
	return rt, nil
}

func (db *personRepository) DeleteRefreshToken(ctx context.Context, personId uuid.UUID, refreshToken string) error {
	return nil
}

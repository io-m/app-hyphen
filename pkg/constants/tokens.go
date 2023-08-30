package constants

import "time"

const (
	REFRESH_TOKEN_KEY        = "refresh_token"
	ACCESS_TOKEN_KEY         = "access_token"
	REFRESH_TOKEN_EXPIRATION = 6 * 30 * 24 * time.Hour
	ACCESS_TOKEN_EXPIRATION  = 5 * time.Minute
)

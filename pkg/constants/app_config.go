package constants

import "time"

type Authenticator string

// Config related constants
const (
	RUNNING_ENV                            = "ENV"
	DEVELOPMENT                            = "DEV"
	PRODUCTION                             = "PROD"
	DEV_CONFIG_FILE          string        = "./infrastructure/dev.env"
	PROD_CONFIG_FILE         string        = "./.env"
	PASETO                   Authenticator = "paseto"
	JWT                      Authenticator = "jwt"
	ACCESS_TOKEN_DURATION    time.Duration = 5 * time.Minute
	REFRESH_TOKEN_DURATION   time.Duration = 100 * 24 * 60 * time.Minute // 100 days
	ACCESS_TOKEN_SECRET_KEY                = "ACCESS_TOKEN_SECRET_KEY"
	REFRESH_TOKEN_SECRET_KEY               = "REFRESH_TOKEN_SECRET_KEY"
)

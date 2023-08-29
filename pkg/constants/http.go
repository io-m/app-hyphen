package constants

type Claims string
type ResponseMapKey string

const (
	CONTENT_TYPE         string         = "Content-Type"
	APPLICATION_JSON     string         = "application/json"
	MAX_SIZE_PAYLOAD     int64          = 1048576 // 1MB of payload is allowed
	ACCESS_TOKEN_HEADER                 = "access-token"
	REFRESH_TOKEN_HEADER                = "refresh-token"
	BEARER_TOKEN                        = "authorization"
	ERROR                ResponseMapKey = "error"
	SUCCESS              ResponseMapKey = "success"
	CLAIMS               Claims         = "claims"
	BASE_ROUTE           string         = "/api/v1"
	BASE_ROUTE_PROTECTED string         = "/api/v1/p"
)

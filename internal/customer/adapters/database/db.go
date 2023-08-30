package customer_db_adapter

import (
	"github.com/arangodb/go-driver"
	"github.com/go-redis/redis/v8"
)

type customerOutgoing struct {
	arango driver.Database
	redis  *redis.Client
}

func NewCustomerOutgoing(arangoDriver driver.Database, redisClient *redis.Client) *customerOutgoing {
	return &customerOutgoing{
		arango: arangoDriver,
		redis:  redisClient,
	}
}

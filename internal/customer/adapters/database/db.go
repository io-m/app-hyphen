package customer_db_adapter

import (
	"github.com/arangodb/go-driver"
	"github.com/go-redis/redis/v8"
)

type customerRepository struct {
	arango driver.Database
	redis  *redis.Client
}

func NewCustomerRepository(arangoDriver driver.Database, redisClient *redis.Client) *customerRepository {
	return &customerRepository{
		arango: arangoDriver,
		redis:  redisClient,
	}
}

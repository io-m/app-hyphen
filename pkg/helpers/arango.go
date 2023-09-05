package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/io-m/app-hyphen/pkg/constants"
)

func CreateArangoConnection() (driver.Database, error) {
	// ARANGO DB SETUP
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{os.Getenv(constants.ARANGO_ENDPOINTS)},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(os.Getenv(constants.ARANGO_USERNAME), os.Getenv(constants.ARANGO_PASSWORD)),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create client %w", err)
	}
	ctx := context.Background()
	customersCollectionExists, _ := c.DatabaseExists(ctx, os.Getenv(constants.ARANGO_DB_NAME))
	if !customersCollectionExists {
		log.Print("Does not exists ...")
		if _, err := c.CreateDatabase(ctx, os.Getenv(constants.ARANGO_DB_NAME), nil); err != nil {
			return nil, fmt.Errorf("failed to create database %s: %w", constants.ARANGO_DB_NAME, err)
		}
	}

	arangoDriver, err := c.Database(ctx, os.Getenv(constants.ARANGO_DB_NAME))
	if err != nil {
		return nil, fmt.Errorf("failed to obtain database %s: %w", os.Getenv(constants.ARANGO_DB_NAME), err)
	}
	return arangoDriver, nil
}

func ReadSingleDocument[T any](ctx context.Context, cursor driver.Cursor) (T, error) {
	var target T
	count := 0
	for cursor.HasMore() {
		_, err := cursor.ReadDocument(ctx, &target)
		if err != nil {
			return target, fmt.Errorf("failed to read document: %w", err)
		}
		count++
	}
	// if count < 1 {
	// 	log.Println("ERROR in COUNT --> ")
	// 	return target, fmt.Errorf("document not found. expected type %v", reflect.TypeOf(target))
	// }
	return target, nil
}

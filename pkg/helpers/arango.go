package helpers

import (
	"context"
	"fmt"
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
		Authentication: driver.BasicAuthentication(constants.ARANGO_USERNAME, os.Getenv(constants.ARANGO_PASSWORD)),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create client %w", err)
	}
	ctx := context.Background()

	// if _, err := c.CreateDatabase(ctx, config.HyphenArangoDBName, nil); err != nil {
	// 	return nil, fmt.Errorf("failed to create database %s: %w", config.HyphenArangoDBName, err)
	// }

	arangoDriver, err := c.Database(ctx, constants.ARANGO_DB_NAME)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain database %s: %w", constants.ARANGO_DB_NAME, err)
	}
	return arangoDriver, nil
}

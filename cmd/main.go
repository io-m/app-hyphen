package main

import (
	"log"
	"net/http"
	"os"

	app "github.com/io-m/app-hyphen/internal"
	hyphen_arango "github.com/io-m/app-hyphen/pkg/arango"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/helpers"
	hyphen_redis "github.com/io-m/app-hyphen/pkg/redis"
	"golang.org/x/exp/slog"
)

func main() {
	helpers.LoadEnv(constants.PROD_CONFIG_FILE)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	arangoDriver, err := hyphen_arango.CreateArangoConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	key, _ := helpers.GenerateRandomKey(16)
	log.Println("RANDOM::: ", key)
	redisClient, err := hyphen_redis.CreateRedisConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Deferring Redis conn closing
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	port := os.Getenv("APP_PORT")
	log.Printf("listening on port: %s............\n", port)
	if err := http.ListenAndServe(":"+port, app.SetAndRun(arangoDriver, redisClient)); err != nil {
		log.Fatalf("Server is down!")
	}
}

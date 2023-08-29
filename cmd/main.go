package main

import (
	"log"
	"net/http"
	"os"

	di "github.com/io-m/app-hyphen/internal"
	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/io-m/app-hyphen/pkg/helpers"
	"golang.org/x/exp/slog"
)

func main() {
	helpers.LoadEnv(constants.PROD_CONFIG_FILE)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	port := os.Getenv("APP_PORT")
	log.Printf("listening on port: %s............\n", port)
	if err := http.ListenAndServe(":"+port, di.SetAndRun()); err != nil {
		log.Fatalf("Server is down!")
	}
}

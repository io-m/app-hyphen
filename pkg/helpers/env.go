package helpers

import (
	"log"

	env "github.com/joho/godotenv"
)

func LoadEnv(path ...string) {
	if err := env.Load(path...); err != nil {
		log.Fatal("loading env variables:", err)
	}
}

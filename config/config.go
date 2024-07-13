package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// C is the global configuration object
var C config

type config struct {
	AppEnv string `envDefault:"dev"`
	DbUrl  string `env:"DB_URL,expand" envDefault:"postgresql://postgres:pass@localhost:5432"`
	Host   string `envDefault:"localhost"`
	Port   string `envDefault:"8080"`
}

func init() {
	if os.Getenv("APP_ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	C = config{}
	opts := env.Options{UseFieldNameByDefault: true}
	if err := env.ParseWithOptions(&C, opts); err != nil {
		log.Fatal(err)
	}
}

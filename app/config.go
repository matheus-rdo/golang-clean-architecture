package app

import (
	"log"

	env "github.com/Netflix/go-env"
	"github.com/subosito/gotenv"
)

// Environment variables
type Environment struct {
	AwsEndpoint   string `env:"AWS_ENDPOINT"`
	AwsRegion     string `env:"AWS_REGION" binding:"required"`
	APIPort       string `env:"API_PORT" binding:"required"`
	BookTableName string `env:"BOOK_TABLE_NAME" binding:"required"`
}

// ENV - output variable
var ENV Environment

func init() {
	gotenv.Load("../.env") // load .env file (if exists)
	if _, err := env.UnmarshalFromEnviron(&ENV); err != nil {
		log.Fatal("Fatal error unmarshalling environment config: ", err)
	}
}

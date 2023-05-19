package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApiUrl   = ""
	ApiPort  = 0
	HashKey  []byte
	BlockKey []byte
)

func LoadConfig() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Fatal("Fail on load environment variables")
	}

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ApiPort = 5000
	}

	ApiUrl = os.Getenv("API_URL")

	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}

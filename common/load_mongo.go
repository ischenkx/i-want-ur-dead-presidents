package common

import (
	"fmt"
	"os"
)

func LoadMongoFromEnv() string {
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	mongoServer := os.Getenv("MONGO_SERVER")

	url := fmt.Sprintf("mongodb://%s:%s@%s:27017/?retryWrites=true&w=majority", user, password, mongoServer)

	return url
}
package core

import (
	"os"
	"sync"
)

var (
	variables *Variables
	once      sync.Once
)

type Variables struct {
	SERVER_PORT    string
	SERVER_HOST    string
	PUBLIC_PATH    string
	PRIVATE_PATH   string
	DB_PATH        string
	PUBLIC_SSSHKEY string
}

func GetVars() *Variables {
	once.Do(func() {
		variables = &Variables{
			SERVER_PORT:    readEnv("SERVER_PORT", "8080"),
			SERVER_HOST:    readEnv("SERVER_HOST", "localhost"),
			PUBLIC_PATH:    readEnv("PUBLIC_PATH", "./data/public"),
			PRIVATE_PATH:   readEnv("PRIVATE_PATH", "./data/private"),
			DB_PATH:        readEnv("DB_PATH", "./data/_database"),
			PUBLIC_SSSHKEY: readEnv("PUBLIC_SSSHKEY", ""),
		}
	})
	return variables
}

func readEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = fallback
	}
	return value
}

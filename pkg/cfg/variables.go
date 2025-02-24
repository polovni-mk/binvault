package cfg

import (
	"os"
	"path/filepath"
	"reflect"
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
	TEMP_PATH      string
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
			TEMP_PATH:      readEnv("TEMP_PATH", "./data/temp"),
			DB_PATH:        readEnv("DB_PATH", "./data/_database"),
			PUBLIC_SSSHKEY: readEnv("PUBLIC_SSSHKEY", ""),
		}
	})
	return variables
}

func GetVar(key string) string {
	vars := GetVars()
	v := reflect.ValueOf(vars)
	field := v.Elem().FieldByName(key)
	return field.String()
}

func GetPath(key string) string {
	path := GetVar(key)
	abs, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return abs
}

func readEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = fallback
	}
	return value
}

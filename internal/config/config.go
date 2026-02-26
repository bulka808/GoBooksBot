package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiId       int
	ApiHash     string
	Phone       string
	BooksChatID int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}
	apiId, err := strconv.Atoi(os.Getenv("API_ID"))
	if err != nil {
		return nil, err
	}
	apiHash := os.Getenv("API_HASH")
	if apiHash == "" {
		return nil, &configError{msg: "API_HASH environment variable not set"}
	}
	phone := os.Getenv("PHONE")
	if phone == "" {
		return nil, &configError{msg: "PHONE environment variable not set"}
	}

	cfg.ApiId = apiId
	cfg.ApiHash = apiHash
	cfg.Phone = phone
	return cfg, nil
}

type configError struct {
	msg string
}

func (e *configError) Error() string {
	return e.msg
}

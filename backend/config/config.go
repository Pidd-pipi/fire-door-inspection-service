package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func Load() Config {
	port := 8080
	if value, err := strconv.Atoi(os.Getenv("PORT")); err == nil && value > 0 && value < 65536 {
		port = value
	}
	return Config{Port: port}
}

func (c Config) Address() string { return fmt.Sprintf(":%d", c.Port) }

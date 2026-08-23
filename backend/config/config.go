package config

import (
	"fmt"
	"os"
	"strconv"
)

const defaultPort = 8080

type Config struct {
	Port int
}

func Load() Config {
	return Config{Port: resolvePort(os.Getenv("PORT"))}
}

// resolvePort parses a PORT value and falls back to the default when it is
// absent, non-numeric, or out of the valid range. A zero or out-of-range port
// is rejected so the server never starts on an unusable address.
func resolvePort(raw string) int {
	if raw == "" {
		return defaultPort
	}
	port, err := strconv.Atoi(raw)
	if err != nil || port < 1 || port > 65535 {
		fmt.Fprintf(os.Stderr, "config: invalid PORT %q, using default %d\n", raw, defaultPort)
		return defaultPort
	}
	return port
}

func (c Config) Address() string { return fmt.Sprintf(":%d", c.Port) }

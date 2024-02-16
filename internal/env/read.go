package env

import (
	"fmt"

	"github.com/caarlos0/env/v10"
)

// Read is a function that parses environment variables into the provided cfg structure.
// The cfg parameter should be a pointer to a struct where each field represents an environment variable.
// The function returns an error if the parsing process fails, wrapping the original error with a message.
func Read(cfg any) error {
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("failed to parse env: %w", err)
	}
	return nil
}

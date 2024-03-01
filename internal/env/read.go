package env

import (
	"fmt"

	"github.com/caarlos0/env/v10"

	"github.com/dsbasko/go-cfg/internal/dflt"
	"github.com/dsbasko/go-cfg/internal/reflect"
)

// Read is a function that parses environment variables into the provided cfg structure.
// The cfg parameter should be a pointer to a struct where each field represents an
// environment variable. The function returns an error if the parsing process fails,
// wrapping the original error with a message.
func Read(structPtr any) error {
	if err := reflect.Validation(structPtr); err != nil {
		return fmt.Errorf("error validating struct: %w", err)
	}

	if err := dflt.Read(structPtr); err != nil {
		return fmt.Errorf("error setting default values: %w", err)
	}

	if err := env.Parse(structPtr); err != nil {
		return fmt.Errorf("failed to parse env: %w", err)
	}

	return nil
}

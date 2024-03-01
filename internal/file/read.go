package file

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"

	"github.com/dsbasko/go-cfg/internal/dflt"
	"github.com/dsbasko/go-cfg/internal/reflect"
)

// Read is a function that parses the content of the file into the provided cfg structure.
// The path parameter should be a string representing the path to the file. The structPtr
// parameter should be a pointer to a struct where each field represents a configuration
// option. The function returns an error if the parsing process fails, wrapping the
// original error with a message.
func Read(path string, structPtr any) error {
	if errValidation := reflect.Validation(structPtr); errValidation != nil {
		return fmt.Errorf("error validating struct: %w", errValidation)
	}

	if errDefault := dflt.Read(structPtr); errDefault != nil {
		return fmt.Errorf("error setting default values: %w", errDefault)
	}

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	switch extension := strings.ToLower(filepath.Ext(path)); extension {
	case ".json":
		if err = parseJSON(file, structPtr); err != nil {
			return fmt.Errorf("failed to parse json: %w", err)
		}
	case ".yaml", ".yml":
		if err = parseYAML(file, structPtr); err != nil {
			return fmt.Errorf("failed to parse yaml: %w", err)
		}
	case ".toml":
		if err = parseTOML(file, structPtr); err != nil {
			return fmt.Errorf("failed to parse toml: %w", err)
		}
	case ".env":
		if err = parseENV(file, structPtr); err != nil {
			return fmt.Errorf("failed to parse env: %w", err)
		}
	}

	return nil
}

// parseJSON is a helper function used by Read to parse the JSON content of the file.
// It takes an io.Reader and a pointer to a struct where each field represents a
// configuration option. The function returns an error if the parsing process fails.
func parseJSON(r io.Reader, structPtr any) error {
	return json.NewDecoder(r).Decode(structPtr)
}

// parseYAML is a helper function used by Read to parse the YAML content of the file.
// It takes an io.Reader and a pointer to a struct where each field represents a
// configuration option. The function returns an error if the parsing process fails.
func parseYAML(r io.Reader, structPtr any) error {
	return yaml.NewDecoder(r).Decode(structPtr)
}

// parseTOML is a helper function used by Read to parse the TOML content of the file.
// It takes an io.Reader and a pointer to a struct where each field represents a
// configuration option. The function returns an error if the parsing process fails.
func parseTOML(r io.Reader, structPtr any) error {
	_, err := toml.NewDecoder(r).Decode(structPtr)
	return err
}

// parseENV is a helper function used by Read to parse the ENV content of the file.
// It takes an io.Reader and a pointer to a struct where each field represents a
// configuration option. The function returns an error if the parsing process fails.
func parseENV(r io.Reader, structPtr any) error {
	dataEnv, err := godotenv.Parse(r)
	if err != nil {
		return fmt.Errorf("failed to parse env: %w", err)
	}

	parsedStruct, err := reflect.ParseTag(structPtr, "env")
	if err != nil {
		return fmt.Errorf("failed to parse tag: %w", err)
	}

	if errReflection := reflect.WriteToStruct(structPtr, func(fieldName string) string {
		if _, ok := parsedStruct[fieldName]; !ok {
			return ""
		}
		return dataEnv[parsedStruct[fieldName].TagValue]
	}); errReflection != nil {
		return errReflection
	}

	return nil
}

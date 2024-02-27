package file

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/dsbasko/go-cfg/pkg/structs"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// Read is a function that reads configuration from a file into the provided cfg structure.
// The path parameter is the path to the configuration file.
// The cfg parameter should be a pointer to a struct where each field represents a configuration option.
// The function returns an error if the reading process fails.
func Read(path string, cfg any) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	switch extension := strings.ToLower(filepath.Ext(path)); extension {
	case ".json":
		if err = parseJSON(file, cfg); err != nil {
			return fmt.Errorf("failed to parse json: %w", err)
		}
	case ".yaml", ".yml":
		if err = parseYAML(file, cfg); err != nil {
			return fmt.Errorf("failed to parse yaml: %w", err)
		}
	case ".toml":
		if err = parseTOML(file, cfg); err != nil {
			return fmt.Errorf("failed to parse toml: %w", err)
		}
	case ".env":
		if err = parseENV(file, cfg); err != nil {
			return fmt.Errorf("failed to parse env: %w", err)
		}
	}

	return nil
}

// parseJSON is a helper function used by Read to parse the JSON content of the file.
// It takes an io.Reader and a pointer to a struct where each field represents a configuration option.
// The function returns an error if the parsing process fails.
func parseJSON(r io.Reader, cfg any) error {
	return json.NewDecoder(r).Decode(cfg)
}

// parseYAML is a helper function used by Read to parse the YAML content of the file.
// It takes an io.Reader and a pointer to a struct where each field represents a configuration option.
// The function returns an error if the parsing process fails.
func parseYAML(r io.Reader, cfg any) error {
	return yaml.NewDecoder(r).Decode(cfg)
}

// parseTOML is a helper function used by Read to parse the TOML content of the file.
// It takes an io.Reader and a pointer to a struct where each field represents a configuration option.
// The function returns an error if the parsing process fails.
func parseTOML(r io.Reader, cfg any) error {
	_, err := toml.NewDecoder(r).Decode(cfg)
	return err
}

// parseENV parses the environment variables from the given io.Reader and populates the provided configuration struct.
// It uses the godotenv package to parse the data.
// If an error occurs during parsing or populating the struct, it returns the error.
func parseENV(r io.Reader, cfg any) error {
	data, err := godotenv.Parse(r)
	if err != nil {
		return err
	}

	if err := structs.FromMap(data, cfg); err != nil {
		return err
	}

	return nil
}

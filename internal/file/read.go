package file

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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

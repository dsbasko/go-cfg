package file

import (
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Read(t *testing.T) {
	type InStructNested struct {
		Field string `json:"field" yaml:"field" toml:"field" env:"NESTED_FIELD"`
	}
	type InStruct struct {
		Field  string         `json:"field" yaml:"field" toml:"field" env:"FIELD"`
		Nested InStructNested `json:"nested" yaml:"nested" toml:"nested"`
	}

	tests := []struct {
		name       string
		path       string
		structPtr  any
		wantStruct *InStruct
		wantError  bool
	}{
		{
			name:      "Happy Path JSON",
			path:      path.Join("tests", "cfg.json"),
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "jsonFieldValue",
				Nested: InStructNested{
					Field: "jsonNestedFieldValue",
				},
			},
			wantError: false,
		},
		{
			name:       "Broken JSON",
			path:       path.Join("tests", "broken.json"),
			structPtr:  &InStruct{},
			wantStruct: &InStruct{},
			wantError:  true,
		},
		{
			name:      "Happy Path YAML",
			path:      path.Join("tests", "cfg.yaml"),
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "yamlFieldValue",
				Nested: InStructNested{
					Field: "yamlNestedFieldValue",
				},
			},
			wantError: false,
		},
		{
			name:       "Broken YAML",
			path:       path.Join("tests", "broken.yaml"),
			structPtr:  &InStruct{},
			wantStruct: &InStruct{},
			wantError:  true,
		},
		{
			name:      "Happy Path TOML",
			path:      path.Join("tests", "cfg.toml"),
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "tomlFieldValue",
				Nested: InStructNested{
					Field: "tomlNestedFieldValue",
				},
			},
			wantError: false,
		},
		{
			name:       "Broken TOML",
			path:       path.Join("tests", "broken.toml"),
			structPtr:  &InStruct{},
			wantStruct: &InStruct{},
			wantError:  true,
		},
		{
			name:      "Happy Path ENV",
			path:      path.Join("tests", "cfg.env"),
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "envFieldValue",
				Nested: InStructNested{
					Field: "envNestedFieldValue",
				},
			},
			wantError: false,
		},
		{
			name:       "Broken ENV",
			path:       path.Join("tests", "broken.env"),
			structPtr:  &InStruct{},
			wantStruct: &InStruct{},
			wantError:  true,
		},
		{
			name:       "Fail Path",
			path:       path.Join("not", "found"),
			structPtr:  &InStruct{},
			wantStruct: &InStruct{},
			wantError:  true,
		},
		{
			name:       "Validation Not Pointer",
			path:       path.Join("tests"),
			structPtr:  InStruct{},
			wantStruct: &InStruct{},
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Read(tt.path, tt.structPtr)
			if err != nil {
				if (err != nil) != tt.wantError {
					t.Errorf("Read() error = %v, wantErr %v", err, tt.wantError)
				}
				return
			}
			assert.EqualValues(t, tt.wantStruct, tt.structPtr)
		})
	}
}

func Test_parseJSON(t *testing.T) {
	type InStructNested struct {
		Field string `json:"field"`
	}
	type InStruct struct {
		Field  string         `json:"field"`
		Nested InStructNested `json:"nested"`
	}

	tests := []struct {
		name       string
		reader     func() io.Reader
		structPtr  any
		wantStruct *InStruct
		wantErr    bool
	}{
		{
			name: "Happy Path",
			reader: func() io.Reader {
				file, err := os.Open(path.Join("tests", "cfg.json"))
				assert.NoError(t, err)
				return file
			},
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "jsonFieldValue",
				Nested: InStructNested{
					Field: "jsonNestedFieldValue",
				},
			},
			wantErr: false,
		},
		{
			name: "Parse Error",
			reader: func() io.Reader {
				return strings.NewReader(`not json`)
			},
			structPtr:  &InStruct{},
			wantStruct: &InStruct{},
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseJSON(tt.reader(), tt.structPtr); (err != nil) != tt.wantErr {
				t.Errorf("parseJSON() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.EqualValues(t, tt.wantStruct, tt.structPtr)
			}
		})
	}
}

func Test_parseYAML(t *testing.T) {
	type InStructNested struct {
		Field string `yaml:"field"`
	}
	type InStruct struct {
		Field  string         `yaml:"field"`
		Nested InStructNested `yaml:"nested"`
	}

	tests := []struct {
		name       string
		reader     func() io.Reader
		structPtr  any
		wantStruct *InStruct
		wantErr    bool
	}{
		{
			name: "Happy Path",
			reader: func() io.Reader {
				file, err := os.Open(path.Join("tests", "cfg.yaml"))
				assert.NoError(t, err)
				return file
			},
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "yamlFieldValue",
				Nested: InStructNested{
					Field: "yamlNestedFieldValue",
				},
			},
			wantErr: false,
		},
		{
			name: "Parse Error",
			reader: func() io.Reader {
				return strings.NewReader(`not yaml`)
			},
			structPtr:  &InStruct{},
			wantStruct: &InStruct{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseYAML(tt.reader(), tt.structPtr); (err != nil) != tt.wantErr {
				t.Errorf("parseYAML() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.EqualValues(t, tt.wantStruct, tt.structPtr)
			}
		})
	}
}

func Test_parseTOML(t *testing.T) {
	type InStructNested struct {
		Field string `toml:"field"`
	}
	type InStruct struct {
		Field  string         `toml:"field"`
		Nested InStructNested `toml:"nested"`
	}

	tests := []struct {
		name       string
		reader     func() io.Reader
		structPtr  any
		wantStruct *InStruct
		wantErr    bool
	}{
		{
			name: "Happy Path",
			reader: func() io.Reader {
				file, err := os.Open(path.Join("tests", "cfg.toml"))
				assert.NoError(t, err)
				return file
			},
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "tomlFieldValue",
				Nested: InStructNested{
					Field: "tomlNestedFieldValue",
				},
			},
			wantErr: false,
		},
		{
			name: "Parse Error",
			reader: func() io.Reader {
				return strings.NewReader(`not toml`)
			},
			structPtr:  &InStruct{},
			wantStruct: &InStruct{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseTOML(tt.reader(), tt.structPtr); (err != nil) != tt.wantErr {
				t.Errorf("parseTOML() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.EqualValues(t, tt.wantStruct, tt.structPtr)
			}
		})
	}
}

func Test_parseENV(t *testing.T) {
	type InStructNested struct {
		Field string `env:"NESTED_FIELD"`
	}
	type InStruct struct {
		Field  string `env:"FIELD"`
		Nested InStructNested
	}

	tests := []struct {
		name       string
		reader     func() io.Reader
		structPtr  any
		wantStruct *InStruct
		wantErr    bool
	}{
		{
			name: "Happy Path",
			reader: func() io.Reader {
				file, err := os.Open(path.Join("tests", "cfg.env"))
				assert.NoError(t, err)
				return file
			},
			structPtr: &InStruct{},
			wantStruct: &InStruct{
				Field: "envFieldValue",
				Nested: InStructNested{
					Field: "envNestedFieldValue",
				},
			},
			wantErr: false,
		},
		{
			name: "Parse Error",
			reader: func() io.Reader {
				return strings.NewReader(`{}`)
			},
			structPtr:  &InStruct{},
			wantStruct: &InStruct{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseENV(tt.reader(), tt.structPtr); (err != nil) != tt.wantErr {
				t.Errorf("parseTOML() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.EqualValues(t, tt.wantStruct, tt.structPtr)
			}
		})
	}
}

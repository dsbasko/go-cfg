# Go Configuration Library

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/dsbasko/go-cfg)

[![GitHub Workflow](https://github.com/dsbasko/go-cfg/actions/workflows/go.yaml/badge.svg?branch=main)](https://github.com/dsbasko/go-cfg/actions/workflows/go.yaml)
[![Test Coverage](https://codecov.io/gh/dsbasko/go-cfg/graph/badge.svg?token=142JTUL2X5)](https://codecov.io/gh/dsbasko/go-cfg)
[![Go Report Card](https://goreportcard.com/badge/github.com/dsbasko/go-cfg)](https://goreportcard.com/report/github.com/dsbasko/go-cfg)
![GitHub Version](https://img.shields.io/github/go-mod/go-version/dsbasko/go-cfg.svg)
![GitHub Tag](https://img.shields.io/github/tag/dsbasko/go-cfg.svg)

This project is a Go library for reading configuration data from various sources such as environment variables, command-line flags, and configuration files. The library provides a unified interface for reading configuration data, making it easier to manage and maintain your application's configuration.  

## Attention
The library uses the [env](https://github.com/caarlos0/env), [pflag](https://github.com/spf13/pflag), [yaml](https://github.com/go-yaml/yaml), [toml](https://github.com/BurntSushi/toml) and [godotenv](https://github.com/joho/godotenv) codebase to work with environment variables and flags. This is a temporary solution, maybe I’ll write my own implementation later. Thanks to the authors of these libraries for the work done!

### Installation
To install the library, use the go get command:
```bash
go get github.com/dsbasko/go-cfg
```

## Usage
The library provides several functions for reading configuration data: 

- `ReadEnv(cfg any) error`: Reads environment variables into the provided `cfg` structure. Each field in the `cfg` structure represents an environment variable.  
- `MustReadEnv(cfg any)`: Similar to `ReadEnv` but panics if the reading process fails.  
- `ReadFlag(cfg any) error`: Reads command-line flags into the provided `cfg` structure. Each field in the `cfg` structure represents a command-line flag.  
- `MustReadFlag(cfg any)`: Similar to `ReadFlag` but panics if the reading process fails.  
- `ReadFile(path string, cfg any) error`: Reads configuration from a file into the provided `cfg` structure. The path parameter is the path to the configuration file. Each field in the `cfg` structure represents a configuration option. Supported file formats include JSON, YAML, TOML and .env.
- `MustReadFile(path string, cfg any)`: Similar to `ReadFile` but panics if the reading process fails.

Here is an example of how to use the library:

```go
package main

import (
	"github.com/dsbasko/go-cfg"
)

type Config struct {
	Mode string `default:"prod" json:"mode" yaml:"mode" s-flag:"m" flag:"mode" env:"MODE" description:"mode of the application (dev|prod)"`
	HTTP struct {
		Host         string `default:"localhost" json:"host" yaml:"host" s-flag:"h" flag:"http-host" env:"HTTP_HOST"`
		Port         int    `default:"3000" json:"port" yaml:"port" s-flag:"p" flag:"http-port" env:"HTTP_PORT"`
		ReadTimeout  int    `json:"read_timeout" yaml:"read-timeout" flag:"http-read-timeout" env:"HTTP_READ_TIMEOUT"`
		WriteTimeout int    `json:"write_timeout" yaml:"write-timeout" flag:"http-write-timeout" env:"HTTP_WRITE_TIMEOUT"`
	} `json:"http" yaml:"http"`
}

func main() {
	cfg := Config{}

	gocfg.MustReadFile("configs/config.yaml", &cfg)
	gocfg.MustReadEnv(&cfg)
	gocfg.MustReadFlag(&cfg)
}
```

Note that you can configure the priority of the configuration. For example, you can first read YAML configs, then environment variables, and finally flags, or vice versa.

## Flags

Run a project with flags: `go run ./cmd/main.go -s="some short flag" --flat=f1 --nested n1`

```go
type config struct {
	WithShort string `s-flag:"s" flag:"with-short" description:"With short flag"`
	Flat      string `flag:"flat" description:"Flat flag"`
	Parent    struct {
		Nested string `flag:"nested" description:"Nested flag"`
	}
}

func main() {
	var cfg config
	if err := gocfg.ReadFlag(&cfg); err != nil {
		log.Panicf("failed to read flag: %v", err)
	}
	// or: gocfg.MustReadFlag(&cfg)

	log.Printf("WithShort: %v\n", cfg.WithShort)
	log.Printf("Flat: %v\n", cfg.Flat)
	log.Printf("Nested: %v\n", cfg.Parent.Nested)
}

// WithShort: some short flag
// Flat: f1
// Nested: n1
```

Struct tags are available for working with flags:
- `default` default value;
- `flag` the name of the flag;
- `s-flag` short name of the flask (1 symbol);
- `description` description of the flag that is displayed when running the `--help` command.

## Environment variables

The `env` structure tag is used for environment variables.
Run a project with environment variables: `FLAT=f1 NESTED=n1 go run ./cmd/main.go`

```go
type config struct {
	Flat   string `env:"FLAT"`
	Parent struct {
		Nested string `env:"NESTED"`
	}
}

func main() {
	var cfg config
	if err := gocfg.ReadEnv(&cfg); err != nil {
		log.Panicf("failed to read env: %v", err)
	}
	// or: gocfg.MustReadEnv(&cfg)

	log.Printf("Flat: %v\n", cfg.Flat)
	log.Printf("Nested: %v\n", cfg.Parent.Nested)
}

// Flat: f1
// Nested: n1
```

Struct tags are available for working with environment variables:
- `default` default value;
- `env` the name of the environment variable.

## Files

Run a project with environment variables: `go run ./cmd/main.go`

```go
type config struct {
	Flat   string `yaml:"flat"`
	Parent struct {
		Nested string `yaml:"nested"`
	}
}

func main() {
	var cfg config
	if err := gocfg.ReadFile(path.Join("cmd", "config.yaml"), &cfg); err != nil {
		log.Panicf("failed to read config.yaml file: %v", err)
	}
	// or: gocfg.MustReadFile(path.Join("cmd", "config.yaml"), &cfg)

	log.Printf("Flat: %v\n", cfg.Flat)
	log.Printf("Nested: %v\n", cfg.Parent.Nested)
}

// Flat: f1
// Nested: n1
```

Structure of the `yaml` file:
```yaml
flat: f1
foo:
  nested: n1
```

Struct tags are available for working with environment variables:
- `default` default value;
- `env` for files of the format `.env`
- `yaml` for files of the format `.yaml` or `.yml`
- `toml` for files of the format `.toml`
- `json` for files of the format `.json`

<br>

---

If you enjoyed this project, I would appreciate it if you could give it a star! If you notice any problems or have any suggestions for improvement, please feel free to create a new issue. Your feedback means a lot to me!

❤️ [Dmitriy Basenko](https://github.com/dsbasko)
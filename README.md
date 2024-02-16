# Go Configuration Library
This project is a Go library for reading configuration data from various sources such as environment variables, command-line flags, and configuration files. The library provides a unified interface for reading configuration data, making it easier to manage and maintain your application's configuration.  

## Attention
The library uses the [env](github.com/caarlos0/env) and [flaggy](github.com/integrii/flaggy) codebase to work with environment variables and flags. This is a temporary solution, maybe I’ll write my own implementation later. Thanks to the authors of these libraries for the work done!

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
- `ReadFile(path string, cfg any) error`: Reads configuration from a file into the provided `cfg` structure. The path parameter is the path to the configuration file. Each field in the `cfg` structure represents a configuration option. Supported file formats include JSON and YAML.
- `MustReadFile(path string, cfg any)`: Similar to `ReadFile` but panics if the reading process fails.

Here is an example of how to use the library:

```go
package main

import (
	"github.com/dsbasko/go-cfg"
)

type Config struct {
	Mode string `json:"mode" yaml:"mode" s-flag:"m" flag:"mode" env:"MODE" description:"mode of the application (dev|prod)"`
	HTTP struct {
		Host         string `json:"host" yaml:"host" s-flag:"h" flag:"http-host" env:"HTTP_HOST"`
		Port         int    `json:"port" yaml:"port" s-flag:"p" flag:"http-port" env:"HTTP_PORT"`
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


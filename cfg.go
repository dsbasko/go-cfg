package gocfg

import (
	"github.com/dsbasko/go-cfg/internal/env"
	"github.com/dsbasko/go-cfg/internal/file"
	"github.com/dsbasko/go-cfg/internal/flag"
)

// ReadEnv is a function that reads environment variables into the provided cfg structure.
// The cfg parameter should be a pointer to a struct where each field represents an environment variable.
// The function returns an error if the reading process fails.
//
// Example of how to use the ReadEnv function:
//
// Define a struct that represents your environment variables. For example:
//
//	type Config struct {
//		Mode string `env:"MODE"`
//		HTTP struct {
//			Host string `env:"HTTP_HOST"`
//			Port int    `env:"HTTP_PORT"`
//		}
//	}
//
// Then, create an instance of your struct and call the ReadEnv function:
//
//	func main() {
//		cfg := &Config{}
//		if err := gocfg.ReadEnv(cfg); err != nil {
//			log.Fatalf("failed to read environment variables: %v", err)
//		}
//
//		fmt.Printf("Mode: %s\n", cfg.Mode)
//		fmt.Printf("HTTP Host: %s\n", cfg.HTTP.Host)
//		fmt.Printf("HTTP Port: %d\n", cfg.HTTP.Port)
//	}
//
// This will read the MODE, REST_HOST and REST_PORT environment variables into the Mode, HTTP.Host and HTTP.Port fields of the cfg variable.
func ReadEnv(cfg any) error {
	return env.Read(cfg)
}

// MustReadEnv is similar to ReadEnv but panics if the reading process fails.
// This function is useful when the absence of environment variables should lead to a program termination.
//
// Example:
//
//	type Config struct {
//		Mode string `env:"MODE"`
//		HTTP struct {
//			Host string `env:"HTTP_HOST"`
//			Port int    `env:"HTTP_PORT"`
//		}
//	}
//
//	func main() {
//		cfg := &Config{}
//		gocfg.MustReadEnv(cfg)
//
//		fmt.Printf("Mode: %s\n", cfg.Mode)
//		fmt.Printf("HTTP Host: %s\n", cfg.HTTP.Host)
//		fmt.Printf("HTTP Port: %d\n", cfg.HTTP.Port)
//	}
//
// This will read the MODE, HTTP_HOST and HTTP_PORT environment variables into the Mode, HTTP.Host and HTTP.Port fields of the cfg variable.
// If any of these environment variables are not set, the program will panic.
func MustReadEnv(cfg any) {
	if err := ReadEnv(cfg); err != nil {
		panic(err)
	}
}

// ReadFlag reads command-line flags into the provided cfg structure.
// The cfg parameter should be a pointer to a struct where each field represents a command-line flag.
// This function returns an error if the parsing process fails.
//
// Example:
//
//	type Config struct {
//		Mode string `flag:"mode" s-flag:"m" description:"Application mode"`
//		HTTP struct {
//			Host string `flag:"http-host" s-flag:"hh" description:"HTTP host"`
//			Port int    `flag:"http-port" s-flag:"hp" description:"HTTP port"`
//		}
//	}
//
//	func main() {
//		cfg := &Config{}
//		if err := gocfg.ReadFlag(cfg); err != nil {
//			log.Fatalf("failed to read command-line flags: %v", err)
//		}
//
//		fmt.Printf("Mode: %s\n", cfg.Mode)
//		fmt.Printf("HTTP Host: %s\n", cfg.HTTP.Host)
//		fmt.Printf("HTTP Port: %d\n", cfg.HTTP.Port)
//	}
//
// This will read the command-line flags --mode, --http-host and --http-port (or -m, -hh and -hp respectively) into the Mode, HTTP.Host and HTTP.Port fields of the cfg variable.
// If any of these flags are not set, the function will return an error.
func ReadFlag(cfg any) error {
	return flag.Read(cfg)
}

// MustReadFlag is similar to ReadFlag but panics if the reading process fails.
// This function is useful when the absence of command-line flags should lead to a program termination.
//
// Example:
//
//	type Config struct {
//		Mode string `flag:"mode" s-flag:"m" description:"Application mode"`
//		HTTP struct {
//			Host string `flag:"http-host" s-flag:"hh" description:"HTTP host"`
//			Port int    `flag:"http-port" s-flag:"hp" description:"HTTP port"`
//		}
//	}
//
//	func main() {
//		cfg := &Config{}
//		gocfg.MustReadFlag(cfg)
//
//		fmt.Printf("Mode: %s\n", cfg.Mode)
//		fmt.Printf("HTTP Host: %s\n", cfg.HTTP.Host)
//		fmt.Printf("HTTP Port: %d\n", cfg.HTTP.Port)
//	}
//
// This will read the command-line flags --mode, --http-host and --http-port (or -m, -hh and -hp respectively) into the Mode, HTTP.Host and HTTP.Port fields of the cfg variable.
// If any of these flags are not set, the program will panic.
func MustReadFlag(cfg any) {
	if err := ReadFlag(cfg); err != nil {
		panic(err)
	}
}

// ReadFile reads configuration from a file into the provided cfg structure.
// The path parameter is the path to the configuration file.
// The cfg parameter should be a pointer to a struct where each field represents a configuration option.
// This function returns an error if the reading process fails.
//
// Example:
//
//	type Config struct {
//		Mode string `yaml:"mode"`
//		HTTP struct {
//			Host string `yaml:"http_host"`
//			Port int    `yaml:"http_port"`
//		}
//	}
//
//	func main() {
//		cfg := &Config{}
//		if err := gocfg.ReadFile("config.yaml", cfg); err != nil {
//			log.Fatalf("failed to read configuration file: %v", err)
//		}
//
//		fmt.Printf("Mode: %s\n", cfg.Mode)
//		fmt.Printf("HTTP Host: %s\n", cfg.HTTP.Host)
//		fmt.Printf("HTTP Port: %d\n", cfg.HTTP.Port)
//	}
//
// This will read the mode, http_host and http_port configuration options from the config.yaml file into the Mode, HTTP.Host and HTTP.Port fields of the cfg variable.
// If any of these configuration options are not set in the file, the function will return an error.
func ReadFile(path string, cfg any) error {
	return file.Read(path, cfg)
}

// MustReadFile is similar to ReadFile but panics if the reading process fails.
// This function is useful when the absence of a configuration file should lead to a program termination.
//
// Example:
//
//	type Config struct {
//		Mode string `yaml:"mode"`
//		HTTP struct {
//			Host string `yaml:"host"`
//			Port int    `yaml:"port"`
//		} `yaml:"http"`
//	}
//
//	func main() {
//		cfg := &Config{}
//		gocfg.MustReadFile("config.yaml", cfg)
//
//		fmt.Printf("Mode: %s\n", cfg.Mode)
//		fmt.Printf("HTTP Host: %s\n", cfg.HTTP.Host)
//		fmt.Printf("HTTP Port: %d\n", cfg.HTTP.Port)
//	}
//
// This will read the mode, http_host and http_port configuration options from the config.yaml file into the Mode, HTTP.Host and HTTP.Port fields of the cfg variable.
// If any of these configuration options are not set in the file, the program will panic.
func MustReadFile(path string, cfg any) {
	if err := ReadFile(path, cfg); err != nil {
		panic(err)
	}
}
